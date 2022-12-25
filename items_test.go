package hackernews

import (
	"context"
	"testing"
)

func TestClient_BestStories(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	stories, err := c.BestStories(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(stories) == 0 {
		t.Fatal("no stories")
	}
}

func TestClient_Item(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	item, err := c.GetItem(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}
	if item.ID != 1 {
		t.Fatal("wrong item")
	}
}

func TestClient_MaxItemID(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	id, err := c.MaxItemID(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Fatal("no id")
	}
}

func TestClient_NewStories(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	stories, err := c.NewStories(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(stories) == 0 {
		t.Fatal("no stories")
	}
}

func TestClient_TopStories(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	stories, err := c.TopStories(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(stories) == 0 {
		t.Fatal("no stories")
	}
}

func TestClient_AskStories(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	stories, err := c.AskStories(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(stories) == 0 {
		t.Fatal("no stories")
	}
}

func TestClient_ShowStories(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	stories, err := c.ShowStories(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(stories) == 0 {
		t.Fatal("no stories")
	}
}

func TestClient_JobStories(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	stories, err := c.JobStories(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(stories) == 0 {
		t.Fatal("no stories")
	}
}
