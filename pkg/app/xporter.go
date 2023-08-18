package app

import (
	"os"
	"xporter/pkg/scraper"
)

func Setup() error {
	username := os.Getenv("X_USERNAME")
	password := os.Getenv("X_PASSWORD")
	scraper.Init(username, password)
}
