package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hermanschaaf/hackernews"
)

// Print the title of the current top story.
func main() {
	ctx := context.Background()
	client := hackernews.NewClient()

	// Get the top stories
	topStories, err := client.TopStories(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print the title of the first story
	story, err := client.GetItem(ctx, topStories[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(story.Title)
}
