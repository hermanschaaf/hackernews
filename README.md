# Hacker News API Go Client

This is an up-to-date Go client for the [Hacker News API](https://github.com/HackerNews/API).

- Supports Go modules
- Supports context.Context
- Supports all documented endpoints
- Fully tested

## Installation

```bash
go get github.com/hermanschaaf/hackernews
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/hermanschaaf/hackernews"
)

func main() {
    client := hackernews.NewClient()
	ctx := context.Background()

    // Get the top 10 stories
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
```