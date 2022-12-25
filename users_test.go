package hackernews

import (
	"context"
	"testing"
)

func TestClient_GetUser(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	user, err := c.GetUser(ctx, "pg")
	if err != nil {
		t.Fatal(err)
	}
	if user.ID != "pg" {
		t.Fatal("wrong user")
	}
}
