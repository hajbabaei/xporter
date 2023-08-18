package scraper

import (
	"fmt"
	"os"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
	scraper := twitterscraper.New()
	username := os.Getenv("X_USERNAME")
	password := os.Getenv("X_PASSWORD")
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
