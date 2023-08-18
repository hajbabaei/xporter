package scraper

import (
	"fmt"
	"xporter/pkg/domain"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func Init(username, password string) (*domain.XScraper, error) {
	scraper := twitterscraper.New()
	err := scraper.Login(username, password)
	if err != nil {
		panic(err)
	}
	tweet, err := scraper.GetTweet("1328684389388185600")
	if err != nil {
		panic(err)
	}
	fmt.Println(tweet.Text)
}
