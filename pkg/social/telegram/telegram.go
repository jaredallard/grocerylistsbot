package telegram

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jaredallard/grocerylistsbot/ent"
	"github.com/jaredallard/grocerylistsbot/ent/useridmapping"
	"github.com/jaredallard/grocerylistsbot/pkg/api"
	"github.com/jaredallard/grocerylistsbot/pkg/social"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

type Provider struct {
	client *tgbotapi.BotAPI
	api    *api.Client
	ctx    context.Context
	cache  *cache.Cache
}

// NewProvider creates a new Telegram message provider
func NewProvider(ctx context.Context, api *api.Client) (*Provider, error) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		return nil, err
	}

	c := cache.New(30*time.Minute, 1*time.Hour)

	return &Provider{
		client: bot,
		ctx:    ctx,
		api:    api,
		cache:  c,
	}, nil
}

func (p *Provider) processUpdate(update tgbotapi.Update, stream chan social.Message) error {
	log.Infof("got update: %v", update)
	if update.Message == nil { // ignore any non-Message Updates
		log.Infof("skipping non-message update")
		return nil
	}

	username := update.Message.From.UserName
	if username == "" {
		username = update.Message.From.FirstName + update.Message.From.LastName
	}
	username = strings.ToLower(username)

	msg := social.Message{
		ChatID:       strconv.Itoa(int(update.Message.Chat.ID)),
		Username:     username,
		UserID:       strconv.Itoa(update.Message.From.ID),
		PlatformName: useridmapping.PlatformTypeTelegram,
		Text:         update.Message.Text,
		Replyer: func(chatId, text string) error {
			chatID, err := strconv.Atoi(chatId)
			if err != nil {
				return err
			}

			log.Infof("[telegram] sending message: %v", strings.ReplaceAll(text, "\n", "\\n"))

			msg := tgbotapi.NewMessage(int64(chatID), text)
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ParseMode = "MarkdownV2"
			_, err = p.client.Send(msg)
			return err
		},
	}

	cacheKey := fmt.Sprintf("%s:%d", useridmapping.PlatformTypeTelegram, update.Message.From.ID)
	v, found := p.cache.Get(cacheKey)

	// check if we didn't find a user
	if !found || v == nil {
		log.Warnf("cache miss for user: %s", cacheKey)
		u, err := p.api.GetUserBySNS(p.ctx, useridmapping.PlatformTypeTelegram, strconv.Itoa(update.Message.From.ID))
		if err != nil {
			msg.Error = err
		} else {
			msg.From = u
		}
		p.cache.Set(cacheKey, u, cache.DefaultExpiration)
	} else { // we found the user in our cache
		msg.From = v.(*ent.User)
	}

	// TODO(jaredallard): cache users
	stream <- msg
	return nil
}

// CreatStream returns a telegram message stream
func (p *Provider) CreateStream() (<-chan social.Message, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := p.client.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}

	stream := make(chan social.Message)

	go func() {
		// close the channel once we're done
		defer close(stream)

		for {
			select {
			case update := <-updates:
				if err := p.processUpdate(update, stream); err != nil {
					log.Errorf("error processing message: %v", err)
				}
			case <-p.ctx.Done():
				log.Warnf("message processor shutdown")
				return
			}
		}
	}()

	return stream, nil
}

func (p *Provider) Send(m *social.Message) error {
	chatID, err := strconv.Atoi(m.ChatID)
	if err != nil {
		return err
	}

	log.Infof("[telegram] sending message: %v", strings.ReplaceAll(m.Text, "\n", "\\n"))

	msg := tgbotapi.NewMessage(int64(chatID), m.Text)
	msg.ParseMode = "MarkdownV2"
	_, err = p.client.Send(msg)
	return err
}
