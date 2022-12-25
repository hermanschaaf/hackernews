package hackernews

import (
	"context"
	"encoding/json"
	"fmt"
)

type User struct {
	ID        string `json:"id"`
	Created   int    `json:"created"`
	Karma     int    `json:"karma"`
	About     string `json:"about"`
	Submitted []int  `json:"submitted"`
}

func (c *Client) GetUser(ctx context.Context, id string) (User, error) {
	path := fmt.Sprintf("/v0/user/%s.json", id)
	b, err := c.get(ctx, path)
	if err != nil {
		return User{}, fmt.Errorf("while getting user: %w", err)
	}
	user := User{}
	err = json.Unmarshal(b, &user)
	return user, err
}
