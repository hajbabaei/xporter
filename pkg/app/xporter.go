package app

import (
	"os"
	"xporter/pkg/domain"
	"xporter/pkg/scraper"
)

func SetupScraper() *domain.XScraper {
	username := os.Getenv("X_USERNAME")
	password := os.Getenv("X_PASSWORD")
	xscraper := scraper.Init(username, password)

	return xscraper
}
