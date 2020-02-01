package api

import "github.com/jaredallard/grocerylistsbot/ent"

type Client struct {
	db *ent.Client
}

func NewClient(c *ent.Client) *Client {
	return &Client{
		db: c,
	}
}
