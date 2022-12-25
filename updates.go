package hackernews

import (
	"context"
	"encoding/json"
	"fmt"
)

type Updates struct {
	Items []int `json:"items"`
	Profiles []string `json:"profiles"`
}

// Updates returns the latest item IDs.
// The updates represent the latest items that have been created, updated, or deleted.
func (c *Client) Updates(ctx context.Context) (*Updates, error) {
	b, err := c.get(ctx, "/v0/updates.json")
	if err != nil {
		return nil, fmt.Errorf("while getting updates: %w", err)
	}
	var updates Updates
	err = json.Unmarshal(b, &updates)
	return &updates, err
}
