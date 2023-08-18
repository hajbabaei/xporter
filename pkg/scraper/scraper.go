package scraper

import (
	"xporter/pkg/domain"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func Init(username, password string) *domain.XScraper {
	scraper := twitterscraper.New()
	err := scraper.Login(username, password)
	if err != nil {
		panic(err)
	}

	return &domain.XScraper{Scraper: scraper}
}
