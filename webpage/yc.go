package webpage

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type YCStory struct {
	Title string
	URL   string
}

func hackerNewsPage(pageNumber int) string {
	return fmt.Sprintf("https://news.ycombinator.com/news?p=%d", pageNumber)
}

func GetYCStories() []YCStory {
	var wg sync.WaitGroup
	storiesChan := make(chan YCStory)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			url := hackerNewsPage(i)

			document, err := goquery.NewDocument(url)
			if err != nil {
				log.Println("Error loading HackerNews webpage number ", i, err)
			}

			document.Find("tr td .title").Each(func(_ int, s *goquery.Selection) {
				href, exists := s.Find("a").Attr("href")
				if exists {
					storiesChan <- YCStory{Title: s.Find("a").Text(), URL: href}
				}
			})
		}()
	}

	stories := make([]YCStory, 0)
	go func() {
		for story := range storiesChan {
			// retrieve values from channel and append to YCStory slice
			stories = append(stories, story)
		}
	}()
	wg.Wait()
	return stories
}

func FilterYCStories(stories []YCStory) []YCStory {
	filteredStories := make([]YCStory, 0)
	for _, story := range stories {
		if strings.Contains(story.Title, "Python") || strings.Contains(story.Title, " go ") {
			filteredStories = append(filteredStories, story)
		}
	}
	return filteredStories
}
