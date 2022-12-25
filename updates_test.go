package hackernews

import (
	"context"
	"testing"
)

func TestClient_GetUpdates(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	updates, err := c.Updates(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(updates.Items) == 0 {
		t.Fatal("no items")
	}
	if len(updates.Profiles) == 0 {
		t.Fatal("no profiles")
	}
}
