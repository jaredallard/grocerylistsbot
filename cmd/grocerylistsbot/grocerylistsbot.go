package main

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jaredallard/grocerylistsbot/ent"
	"github.com/jaredallard/grocerylistsbot/ent/groceryitem"
	"github.com/jaredallard/grocerylistsbot/ent/migrate"
	"github.com/jaredallard/grocerylistsbot/pkg/api"
	"github.com/jaredallard/grocerylistsbot/pkg/social"
	"github.com/jaredallard/grocerylistsbot/pkg/social/telegram"
	log "github.com/sirupsen/logrus"

	// database driver for ent
	_ "github.com/lib/pq"
)

var helpText = `Hi\! I'm @grocerylistsbot\. I can help you manage your grocery lists\. When you first talked to me, I created you an account\.

If you'd like to create a grocery list, tap/click or type /newlist\. There is no limit to how many lists you can have\.
If you want to create a list with multiple users, then type /newlist username\.\.\. supplying as many users as you want\.
Once you've created a list, you can add an item by sending it to me\! Try it out now :\)

If you'd like to list the items in your current "active" list, tap/click or type /items\.
If you'd like to switch the list that you're currently viewing, type /switch ID\. You can view lists with /lists\.


If you run into any issues, please feel free to message @jaredallard\.`

// notifyUsers of events on Telgram only
func notifyUsers(ctx context.Context, c *api.Client, p *telegram.Provider, msg *social.Message, users []*ent.User, action string) {
	for _, user := range users {
		id, err := c.GetUserSNSID(ctx, user, msg.PlatformName)
		if err != nil {
			log.Warnf("failed to notify user: %s", user.String())
			continue
		}

		if err := p.Send(&social.Message{
			ChatID: id.PlatformID,
			Text:   action,
		}); err != nil {
			log.Warnf("failed to notify user: %s", user.String())
			continue
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	client, err := ent.Open("postgres", "sslmode=disable user=postgres dbname=ent")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: ", err)
	}
	defer client.Close()

	// run the auto migration tool.
	if err := client.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Info("successfully ran migrations")

	c := api.NewClient(client)
	log.Info("starting message processor(s)")
	tp, err := telegram.NewProvider(ctx, c)
	if err != nil {
		log.Fatalf("failed to create telegram provider: %v", err)
	}

	stream, err := tp.CreateStream()
	if err != nil {
		log.Fatalf("failed to create telegram message steam: %v", err)
	}

	log.Infof("started message processor(s)")
	for msg := range stream {
		log.Infof("got message: %s", msg.String())

		// handle errors
		if ent.IsNotFound(msg.Error) {
			err := c.CreateUser(ctx, ent.UserIDMapping{
				PlatformID:   msg.UserID,
				PlatformType: msg.PlatformName,
			}, ent.User{
				Name: msg.Username,
			})
			if err != nil {
				log.Errorf("failed to create a user: %v", err)
				if err := msg.Reply("I failed to create you account\\.\\.\\. Please try later!"); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
					continue
				}
			}

			if err := msg.Reply("Hi! I've created you an account"); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
				continue
			}

			if err := msg.Reply(helpText); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
				continue
			}
		} else if msg.Error != nil {
			if err := msg.Reply("I couldn't process that right now\\. Please try again later, or let @jaredallard know\\!"); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
			}
			continue
		}

		tokens := strings.Split(msg.Text, " ")
		cmd, args := tokens[0], tokens[1:]
		if cmd == "/newlist" {
			users := make([]*ent.User, 0)
			for _, username := range args {
				// skip ourself, we add them in later automatically
				if username == msg.Username {
					continue
				}

				user, err := c.GetUserByUsername(ctx, username)
				if err != nil {
					if err := msg.Reply(fmt.Sprintf("Failed to find user %s", username)); err != nil {
						log.Errorf("failed to send user a reply: %v", err)
					}
					break
				}

				users = append(users, user)
			}

			users = append(users, msg.From)

			list, err := c.CreateGroceryList(ctx, users)
			if err != nil {
				if err := msg.Reply(fmt.Sprintf("Failed to create list: %v", err)); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			msg.From.Edges.ActiveList = list
			if err := c.UpdateUser(ctx, msg.From); err != nil {
				log.Warnf("failed to set active list on user: %v", err)
			}

			if err := msg.Reply("Created a new grocery list"); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
				continue
			}
		} else if cmd == "/help" {
			if err := msg.Reply(helpText); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
				continue
			}
		} else if cmd == "/lists" || cmd == "/list" {
			lists, err := c.FindAllListsForUser(ctx, msg.From)
			if err != nil {
				if err := msg.Reply(fmt.Sprintf("failed to list your grocerylists: %v", err)); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			replyBuf := bytes.NewBuffer([]byte{})
			fmt.Fprintf(replyBuf, "Grocery Lists\n\n")
			for _, list := range lists {
				fmt.Fprintf(replyBuf, "ID: %d\\) %s\n", list.ID, list.Name)
			}
			fmt.Fprintf(replyBuf, "\nSwitch to this list via /switch ID\n")
			fmt.Fprintf(replyBuf, "See items in a list with /items ID")

			if err := msg.Reply(replyBuf.String()); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
				continue
			}
		} else if cmd == "/items" {
			if msg.From.Edges.ActiveList == nil && len(args) == 0 {
				if err := msg.Reply("Please provide an ID, I don't have an active list for you\\!"); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
					continue
				}
				continue
			}

			list := msg.From.Edges.ActiveList
			if len(args) != 0 {
				id, err := strconv.Atoi(args[1])
				if err != nil {
					if err := msg.Reply(fmt.Sprintf("Failed to parse ID: %v", err)); err != nil {
						log.Errorf("failed to send user a reply: %v", err)
						continue
					}
					continue
				}

				list, err = c.GetGroceryListByID(ctx, &ent.GroceryList{ID: id})
				if err != nil {
					if err := msg.Reply(fmt.Sprintf("failed to find grocery list with that id: %v", err)); err != nil {
						log.Errorf("failed to send user a reply: %v", err)
						continue
					}
					continue
				}
			}

			lists, err := c.ListGroceryItems(ctx, list)
			if err != nil {
				if err := msg.Reply(fmt.Sprintf("failed to list your items: %v", err)); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			replyBuf := bytes.NewBuffer([]byte{})
			fmt.Fprintf(replyBuf, "Grocery List '%s' has %d Item\\(s\\)\n\n", list.Name, len(lists))
			for _, item := range lists {
				if item.Status == groceryitem.StatusPurchased {
					fmt.Fprintf(replyBuf, "~")
				}
				fmt.Fprintf(replyBuf, "ID %d\\) *%s*", item.ID, item.Name)
				if item.Status == groceryitem.StatusPurchased {
					fmt.Fprintf(replyBuf, "~")
				}
				fmt.Fprintf(replyBuf, "\n")
			}
			fmt.Fprintf(replyBuf, "\nRemove an item with /remove ID\n")
			fmt.Fprintf(replyBuf, "Mark an item as purchased with /purchased ID")

			if err := msg.Reply(replyBuf.String()); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
				continue
			}
		} else if cmd == "/addmembers" {
			if len(args) == 0 {
				if err := msg.Reply("No usernames provided, please supply them like so: /addmembers username1 username2\\.\\.\\."); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
					continue
				}
				continue
			}

			l := msg.From.Edges.ActiveList
			if l == nil {
				if err := msg.Reply("Please switch to a list first, you can view your lists /lists and then /switch ID"); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
					continue
				}
				continue
			}

			members, err := l.QueryMembers().All(ctx)
			if err != nil {
				log.Error("failed to query members from edge: %v", err)
				if err := msg.Reply("An error occurred"); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}
			l.Edges.Members = members

			existingUsers := make(map[int]bool)
			for _, member := range l.Edges.Members {
				existingUsers[member.ID] = true
			}

			for _, username := range args {
				// skip ourself
				if username == msg.Username {
					continue
				}

				user, err := c.GetUserByUsername(ctx, username)
				if err != nil {
					if err := msg.Reply(fmt.Sprintf("Failed to find user %s", username)); err != nil {
						log.Errorf("failed to send user a reply: %v", err)
					}
					break
				}

				if existingUsers[user.ID] {
					continue
				}

				l.Edges.Members = append(l.Edges.Members, user)
				existingUsers[user.ID] = true
			}

			err = c.UpdateGroceryList(ctx, l)
			if err != nil {
				log.Errorf("failed to update grocery list: %v", err)
				if err := msg.Reply("An error occured"); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}
			if err := msg.Reply("Added members to list\\."); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
			}
		} else if cmd == "/updatename" || cmd == "/setname" {
			if len(args) == 0 {
				if err := msg.Reply("Invalid input, expected /updatename NAME"); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			l := msg.From.Edges.ActiveList
			if l == nil {
				if err := msg.Reply("Please switch to a list first, you can view your lists /lists and then /switch ID"); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			l.Name = strings.Join(args, " ")

			err := c.UpdateGroceryList(ctx, l)
			if err != nil {
				log.Errorf("failed to update name of grocery lists %v: %v", l.ID, err)
				if err := msg.Reply(fmt.Sprintf("Failed to update name: %v", err)); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
			}

			if err := msg.Reply("Updated list name\\."); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
			}
		} else if cmd == "/switch" {
			if len(args) != 1 {
				if err := msg.Reply("Invalid input, expected /switch ID"); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			id, err := strconv.Atoi(args[0])
			if err != nil {
				if err := msg.Reply("Invalid input, expected a numeric ID\\."); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			list := &ent.GroceryList{ID: id}
			l, err := c.GetGroceryListByID(ctx, list)
			if err != nil {
				log.Errorf("failed to find grocery list %v: %v", args[0], err)
				if err := msg.Reply(fmt.Sprintf("Grocery list not found: %v", err)); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			msg.From.Edges.ActiveList = l

			err = c.UpdateUser(ctx, msg.From)
			if err != nil {
				log.Errorf("failed to update active grocery list %v: %v", args[0], err)
				if err := msg.Reply("I'm sorry, an error occured! Please try again later\\."); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			if err := msg.Reply("Switched your active list\\."); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
			}
		} else if cmd == "/purchased" {
			if len(args) != 1 {
				if err := msg.Reply("Invalid input, expected /purchased ID"); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			l := msg.From.Edges.ActiveList
			if l == nil {
				if err := msg.Reply("Please switch to a list first, you can view your lists /lists and then /switch ID"); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			id, err := strconv.Atoi(args[0])
			if err != nil {
				if err := msg.Reply("Invalid input, expected a numeric ID\\."); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			err = c.UpdateGroceryItem(ctx, l, &ent.GroceryItem{ID: id, Status: groceryitem.StatusPurchased})
			if err != nil {
				log.Errorf("failed to mark item as purchased %d: %v", id, err)
				if err := msg.Reply("I'm sorry, an error occured! Please try again later\\."); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			if err := msg.Reply("Marked item as purchased"); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
			}
		} else if cmd == "/remove" {
			if len(args) != 1 {
				if err := msg.Reply("Invalid input, expected /remove ID"); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			l := msg.From.Edges.ActiveList
			if l == nil {
				if err := msg.Reply("Please switch to a list first, you can view your lists /lists and then /switch ID"); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			id, err := strconv.Atoi(args[0])
			if err != nil {
				if err := msg.Reply("Invalid input, expected a numeric ID\\."); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			err = c.DeleteGroceryItem(ctx, l, &ent.GroceryItem{ID: id})
			if err != nil {
				log.Errorf("failed to mark item as purchased %d: %v", id, err)
				if err := msg.Reply("I'm sorry, an error occured! Please try again later\\."); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}
			if err := msg.Reply("Removed Item"); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
			}
		} else if strings.Contains(cmd, "/") {
			if err := msg.Reply(fmt.Sprintf("Unknown command '%s'\\. Please click/type /help for more information\\.", cmd)); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
				continue
			}
		} else {
			l := msg.From.Edges.ActiveList
			if l == nil {
				if err := msg.Reply("Please switch to a list first, you can view your lists /lists and then /switch ID"); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			item := strings.ReplaceAll(msg.Text, "!", "")
			_, err := c.CreateGroceryItem(ctx, &ent.GroceryList{ID: l.ID}, 0, item)
			if err != nil {
				if err := msg.Reply(fmt.Sprintf("failed to add item to your list: %v", err)); err != nil {
					log.Errorf("failed to send user a reply: %v", err)
				}
				continue
			}

			members, err := l.QueryMembers().All(ctx)
			if err != nil {
				log.Warnf("failed to get members for notification: %v", err)
			} else {
				newMembers := make([]*ent.User, 0)
				for _, member := range members {
					if member.ID == msg.From.ID {
						continue
					}

					newMembers = append(newMembers, member)
				}

				notifyUsers(ctx, c, tp, &msg, newMembers, fmt.Sprintf("*%s* added '%s' to the grocery list '%s'", msg.From.Name, item, l.Name))
			}

			if err := msg.Reply(fmt.Sprintf("Added '%s' to your grocery list", msg.Text)); err != nil {
				log.Errorf("failed to send user a reply: %v", err)
				continue
			}
		}
	}
	// TODO(jaredallard): switch between grocery lists with a /switch command?

	cancel()
}
