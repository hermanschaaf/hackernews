package hackernews

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

type Item struct {
	ID          int    `json:"id"`
	Deleted     bool   `json:"deleted"`
	Type        string `json:"type"`
	By          string `json:"by"`
	Time        int    `json:"time"`
	Text        string `json:"text"`
	Dead        bool   `json:"dead"`
	Parent      int    `json:"parent"`
	Kids        []int  `json:"kids"`
	URL         string `json:"url"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Parts       []int  `json:"parts"`
	Descendants int    `json:"descendants"`
}

func (c *Client) GetItem(ctx context.Context, id int) (Item, error) {
	path := fmt.Sprintf("/v0/item/%d.json", id)
	b, err := c.get(ctx, path)
	if err != nil {
		return Item{}, fmt.Errorf("while getting item: %w", err)
	}
	item := Item{}
	err = json.Unmarshal(b, &item)
	return item, err
}

func (c *Client) MaxItemID(ctx context.Context) (int, error) {
	b, err := c.get(ctx, "/v0/maxitem.json")
	if err != nil {
		return 0, fmt.Errorf("while getting max item: %w", err)
	}
	return strconv.Atoi(string(b))
}

// TopStories returns up to 500 of the latest top stories.
func (c *Client) TopStories(ctx context.Context) ([]int, error) {
	b, err := c.get(ctx, "/v0/topstories.json")
	if err != nil {
		return nil, fmt.Errorf("while getting top stories: %w", err)
	}
	var stories []int
	err = json.Unmarshal(b, &stories)
	return stories, err
}

// NewStories returns up to 500 of the newest stories.
func (c *Client) NewStories(ctx context.Context) ([]int, error) {
	b, err := c.get(ctx, "/v0/newstories.json")
	if err != nil {
		return nil, fmt.Errorf("while getting new stories: %w", err)
	}
	var stories []int
	err = json.Unmarshal(b, &stories)
	return stories, err
}

// BestStories returns up to 500 of the best stories.
func (c *Client) BestStories(ctx context.Context) ([]int, error) {
	b, err := c.get(ctx, "/v0/beststories.json")
	if err != nil {
		return nil, fmt.Errorf("while getting best stories: %w", err)
	}
	var stories []int
	err = json.Unmarshal(b, &stories)
	return stories, err
}

// AskStories returns up to 200 of the latest ask stories.
func (c *Client) AskStories(ctx context.Context) ([]int, error) {
	b, err := c.get(ctx, "/v0/askstories.json")
	if err != nil {
		return nil, fmt.Errorf("while getting ask stories: %w", err)
	}
	var stories []int
	err = json.Unmarshal(b, &stories)
	return stories, err
}

// ShowStories returns up to 200 of the latest show stories.
func (c *Client) ShowStories(ctx context.Context) ([]int, error) {
	b, err := c.get(ctx,"/v0/showstories.json")
	if err != nil {
		return nil, fmt.Errorf("while getting show stories: %w", err)
	}
	var stories []int
	err = json.Unmarshal(b, &stories)
	return stories, err
}

// JobStories returns up to 200 of the latest job stories.
func (c *Client) JobStories(ctx context.Context) ([]int, error) {
	b, err := c.get(ctx, "/v0/jobstories.json")
	if err != nil {
		return nil, fmt.Errorf("while getting job stories: %w", err)
	}
	var stories []int
	err = json.Unmarshal(b, &stories)
	return stories, err
}
