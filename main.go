package main

import (
	"log"
	"strings"

	"github.com/brenol/annoying-diego/webpage"
)

type Post struct {
	Title string
	URL   string
}

func filterByTitle(title string) bool {
	return strings.Contains(title, "golang") || strings.Contains(title, " go ") || strings.HasSuffix(title, " Go") || strings.Contains(title, "python")
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
			// we need to ignore x-posts and also self posts (posts made by users asking questions)
			if filterByTitle(p.Data.Title) && !strings.Contains(p.Data.Title, "x-post from r/") && !p.Data.IsSelf {
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
	if len(posts) == 0 {
		log.Println("No News about Go found :(")
		return
	}
	sendMail(generateEmailBody(posts))
}
