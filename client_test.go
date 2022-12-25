package hackernews_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hermanschaaf/hackernews"
)

func TestNewClient(t *testing.T) {
	s := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id": "pg", "karma": 10000}`))
	}))
	c := hackernews.NewClient(
		hackernews.WithBaseURL(s.URL),
		hackernews.WithHTTPClient(s.Client()),
	)
	if c == nil {
		t.Fatal("client is nil")
	}
	ctx := context.Background()
	u, err := c.GetUser(ctx, "pg")
	if err != nil {
		t.Fatal(err)
	}
	if u.ID != "pg" {
		t.Fatalf("wrong user. Got %v, want %v", u.ID, "pg")
	}
}

func ExampleClient_GetUser() {
	client := hackernews.NewClient()

	// Get the karma for user "pg" and print true if it's greater than 10000.
	ctx := context.Background()
	user, err := client.GetUser(ctx, "pg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.Karma > 10000)
	// Output: true
}

func ExampleClient_Item() {
	client := hackernews.NewClient()

	// Print the title of a specific story.
	ctx := context.Background()
	story, err := client.GetItem(ctx,34121082)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(story.Title)
	// Output: Merry Christmas, HN
}
