package main

import (
	"fmt"
	"strings"

	"github.com/brenol/annoying-diego/webpage"
)

type Post struct {
	Title string
	URL   string
}

func filterByTitle(title string) bool {
	return strings.Contains(title, "golang") || strings.Contains(title, " go ")
}

func Filter(stories []webpage.YCStory, redditPosts []webpage.RedditPost) []Post {
	filteredPosts := make([]Post, 0)
	for _, story := range stories {
		if filterByTitle(story.Title) {
			filteredPosts = append(filteredPosts, Post{story.Title, story.URL})
		}
	}

	for _, post := range redditPosts {
		for _, p := range post.Data.Children {
			if filterByTitle(p.Data.Title) {
				filteredPosts = append(filteredPosts, Post{p.Data.Title, p.Data.URL})
			}
		}
	}

	return filteredPosts
}

func main() {
	posts := Filter(
		webpage.GetYCStories(),
		webpage.GetPostsFromSubreddits([]string{"programming", "golang"}),
	)
	fmt.Println("Posts about golang -> ", posts)
}
