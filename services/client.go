package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andrewto30/Reddit-Monitor.git/model"
)

func FetchPosts(subreddit string, limit int) ([]model.RedditChildData, error) {
	url := fmt.Sprintf("https://www.reddit.com/r/%s/new.json?limit=%d", subreddit, limit)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var redditListing model.RedditListing
	if err := json.NewDecoder(resp.Body).Decode(&redditListing); err != nil {
		log.Fatal(err)
	}

	var posts []model.RedditChildData
	for _, child := range redditListing.Data.Children {
		posts = append(posts, child.Data) // Corrected access to child.Data
	}
	return posts, nil
}
