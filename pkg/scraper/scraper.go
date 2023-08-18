package scraper

import (
	twitterscraper "github.com/n0madic/twitter-scraper"
)

type XScraper struct {
	Scraper *twitterscraper.Scraper
}

func Init(username, password string) *XScraper {
	scraper := twitterscraper.New()
	err := scraper.Login(username, password)
	if err != nil {
		panic(err)
	}

	return &XScraper{Scraper: scraper}
}

func (xs *XScraper) GetPost(postId string) *twitterscraper.Tweet {
	tweet, err := xs.Scraper.GetTweet(postId)
	if err != nil {
		panic(err)
	}

	return tweet
}
