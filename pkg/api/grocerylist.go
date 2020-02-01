package api

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jaredallard/grocerylistsbot/ent"
	"github.com/jaredallard/grocerylistsbot/ent/groceryitem"
	"github.com/jaredallard/grocerylistsbot/ent/grocerylist"
	"github.com/jaredallard/grocerylistsbot/ent/predicate"
	"github.com/jaredallard/grocerylistsbot/ent/user"
	log "github.com/sirupsen/logrus"
)

type GroceryListError error

var (
	ErrGroceryListNotFound GroceryListError = errors.New("Failed to find grocery list")
	ErrGroceryListExists   GroceryListError = errors.New("Grocery List already exists for the specified users")
)

// usersToUserIDPredicates is a helper to map user id selects to be predicates
// maintains the order of passed in users
func usersToUserIDPredicates(users []*ent.User) []predicate.User {
	predicates := make([]predicate.User, len(users))
	for i, u := range users {
		predicates[i] = user.ID(u.ID)
	}
	return predicates
}

// CreateGroceryList creates a new grocery list
func (c *Client) CreateGroceryList(ctx context.Context, members []*ent.User) (*ent.GroceryList, error) {
	if len(members) == 0 {
		return nil, fmt.Errorf("members cannot be empty")
	}

	if list, err := c.FindGroceryList(ctx, members); list == nil && err == nil {
		return nil, ErrGroceryListExists
	}

	return c.db.GroceryList.Create().AddMembers(members...).Save(ctx)
}

// GetGroceryListByID returns a grocery list by ID
func (c *Client) GetGroceryListByID(ctx context.Context, l *ent.GroceryList) (*ent.GroceryList, error) {
	list, err := c.db.GroceryList.Query().Where(grocerylist.ID(l.ID)).First(ctx)
	if ent.IsNotFound(err) {
		return nil, ErrGroceryListNotFound
	} else if err != nil {
		return nil, err
	}

	// for some reason WithMembers() isn't working, so we have to "inject" this manually. Not optimal at all.
	members, err := list.QueryMembers().All(ctx)
	if err != nil {
		return nil, err
	}
	list.Edges.Members = members

	return list, nil
}

// CreateGroceryItem creates a new item in a grocery list
func (c *Client) CreateGroceryItem(ctx context.Context, list *ent.GroceryList, price float64, name string) (*ent.GroceryItem, error) {
	return c.db.GroceryItem.Create().SetName(name).SetPrice(price).SetGrocerylist(list).Save(ctx)
}

// ListGroceryItems returns the groceryitems in a given grocerylist
func (c *Client) ListGroceryItems(ctx context.Context, l *ent.GroceryList) ([]*ent.GroceryItem, error) {
	return c.db.GroceryItem.Query().Where(
		groceryitem.HasGrocerylistWith(grocerylist.ID(l.ID)),
	).All(ctx)
}

// UpdateGroceryList updates a grocery list
func (c *Client) UpdateGroceryList(ctx context.Context, l *ent.GroceryList) error {
	list, err := c.GetGroceryListByID(ctx, l)
	if err != nil {
		return err
	}

	q := c.db.GroceryList.Update().Where(grocerylist.ID(l.ID))

	if l.Name != "" {
		q = q.SetName(l.Name)
	}

	if l.Edges.Members != nil {
		// first we make a map of the members we want
		hm := make(map[int]bool)
		for _, member := range l.Edges.Members {
			hm[member.ID] = true
		}

		// now we see which ones aren't in our new list of members, we remove those
		listHm := make(map[int]bool)

		log.Infof("list has %d members already", len(list.Edges.Members))
		for _, member := range list.Edges.Members {
			if !hm[member.ID] {
				log.Infof("removing user %v", member.ID)
				q = q.RemoveMemberIDs(member.ID)
			} else {
				listHm[member.ID] = true
			}
		}

		// now we check which ones are already in the members, we skip the ones that are
		// and we add the ones that aren't
		for _, member := range l.Edges.Members {
			if !listHm[member.ID] {
				log.Infof("adding user %d", member.ID)
				q = q.AddMemberIDs(member.ID)
			}
		}
	}

	_, err = q.SetModifiedAt(time.Now()).Save(ctx)
	return err
}

// UpdateGroceryItem updates a grocery item in a list
func (c *Client) UpdateGroceryItem(ctx context.Context, l *ent.GroceryList, i *ent.GroceryItem) error {
	q := c.db.GroceryItem.Update().Where(
		groceryitem.HasGrocerylistWith(grocerylist.ID(l.ID)),
		groceryitem.ID(i.ID),
	)

	if i.Name != "" {
		q = q.SetName(i.Name)
	}

	// this isn't optimal at all
	if i.Price != 0.0 {
		q = q.SetPrice(i.Price)
	}

	if i.Status != "" {
		q = q.SetStatus(i.Status)
	}

	modified, err := q.SetModifiedAt(time.Now()).Save(ctx)
	if modified != 1 && err == nil {
		return fmt.Errorf("no row modified")
	}

	return err
}

// DeleteGroceryItem deletes a grocery item
func (c *Client) DeleteGroceryItem(ctx context.Context, l *ent.GroceryList, i *ent.GroceryItem) error {
	_, err := c.db.GroceryItem.Delete().Where(
		groceryitem.HasGrocerylistWith(grocerylist.ID(l.ID)),
		groceryitem.ID(i.ID),
	).Exec(ctx)
	return err
}

// FindGroceryList finds a grocery list between a group of people
func (c *Client) FindGroceryList(ctx context.Context, users []*ent.User) (*ent.GroceryList, error) {
	return c.db.GroceryList.Query().Where(
		grocerylist.HasMembersWith(usersToUserIDPredicates(users)...),
	).First(ctx)
}

// FindAllListsForUser returns all grocery lists a user is apart of
func (c *Client) FindAllListsForUser(ctx context.Context, user *ent.User) ([]*ent.GroceryList, error) {
	return c.db.GroceryList.Query().Where(
		grocerylist.HasMembersWith(usersToUserIDPredicates([]*ent.User{user})...),
	).WithMembers().All(ctx)
}
