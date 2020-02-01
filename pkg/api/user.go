package api

import (
	"context"
	"time"

	"github.com/jaredallard/grocerylistsbot/ent"
	"github.com/jaredallard/grocerylistsbot/ent/user"
	"github.com/jaredallard/grocerylistsbot/ent/useridmapping"
)

// func levenshtein(s string, field string) *sql.Predicate {
// 	return sql.P().append(func(b *sql.Builder) {
// 		b.WriteString(" levenshtein(")
// 		b.Ident(field).WriteString(",")
// 		b.Arg(s).WriteString(") ")
// 	})
// }

// // GetUserByUsernameFuzzy fuzzy searches for a user by username
// func (c *Client) GetUserbyUsernameFuzzy(ctx context.Context, username string) (*ent.User, error) {
// 	return c.db.User.Query().Where(predicate.User(func(s *sql.Selector) {
// 		s.Where(levenshtein(username, ""))
// 	}))
// }

// GetUserBySNS returns a user object based on the SNS platform and ID
func (c *Client) GetUserBySNS(ctx context.Context, platformType useridmapping.PlatformType, id string) (*ent.User, error) {
	return c.db.UserIDMapping.Query().Where(
		useridmapping.And(
			useridmapping.PlatformTypeEQ(platformType),
			useridmapping.PlatformIDEQ(id),
		),
	).QueryUser().WithActiveList().First(ctx)
}

// GetUserSNSID returns a user's SNS ID
func (c *Client) GetUserSNSID(ctx context.Context, u *ent.User, platformType useridmapping.PlatformType) (*ent.UserIDMapping, error) {
	return c.db.UserIDMapping.Query().Where(
		useridmapping.And(
			useridmapping.PlatformTypeEQ(platformType),
			useridmapping.HasUserWith(user.ID(u.ID)),
		),
	).First(ctx)
}

// GetUserByUsername returns a user by their username
func (c *Client) GetUserByUsername(ctx context.Context, username string) (*ent.User, error) {
	return c.db.User.Query().Where(user.Name(username)).Only(ctx)
}

// CreateUser creates a new user in the database
func (c *Client) CreateUser(ctx context.Context, id ent.UserIDMapping, u ent.User) error {
	tx, err := c.db.Tx(ctx)
	if err != nil {
		return err
	}

	user, err := tx.User.Create().SetName(u.Name).Save(ctx)
	if err != nil {
		return err
	}

	_, err = tx.UserIDMapping.Create().SetUserID(user.ID).
		SetPlatformID(id.PlatformID).
		SetPlatformType(id.PlatformType).Save(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// UpdateUser updates a user
func (c *Client) UpdateUser(ctx context.Context, u *ent.User) error {
	q := c.db.User.Update().Where(
		user.ID(u.ID),
	)

	if u.Name != "" {
		q = q.SetName(u.Name)
	}

	if u.Edges.ActiveList != nil {
		q = q.SetActiveList(u.Edges.ActiveList)
	}

	_, err := q.SetModifiedAt(time.Now()).Save(ctx)
	return err
}
