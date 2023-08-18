package domain

import twitterscraper "github.com/n0madic/twitter-scraper"

type XPost struct {
	Id      string `json:"id" bson:"id"`
	Url     string `json:"url" bson:"url"`
	Profile string `json:"profile" bson:"profile"`
	Content string `json:"Content" bson:"Content"`
}

type XPorterService interface {
	SetupScraper() *XScraper
	GetPost(postId string) (*XPost, error)
}

type XScraper struct {
	Scraper *twitterscraper.Scraper
}

type ScraperService interface {
	Init(username, password string) *XScraper
	GetPost(postId string) (*XPost, error)
}
