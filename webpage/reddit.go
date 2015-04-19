package webpage

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RedditPost struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			}
		}
	}
}

func subredditURL(subreddit string) string {
	return fmt.Sprintf("http://www.reddit.com/r/%s.json", subreddit)
}

func GetPostsFromSubreddits(subreddits []string) []RedditPost {
	posts := make([]RedditPost, 0)
	for _, subreddit := range subreddits {
		resp, err := http.Get(subredditURL(subreddit))
		if err != nil {
			log.Printf("Error loading reddit webpage %s", subreddit)
		}

		var rp RedditPost
		err = json.NewDecoder(resp.Body).Decode(&rp)
		if err != nil {
			log.Printf("Error unmarshalling response", err)
		}
		posts = append(posts, rp)
	}
	return posts
}
