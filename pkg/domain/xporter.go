package domain

import "xporter/pkg/scraper"

type XPost struct {
	Id      string `json:"id" bson:"id"`
	Url     string `json:"url" bson:"url"`
	Profile string `json:"profile" bson:"profile"`
	Content string `json:"Content" bson:"Content"`
}

type XPorterService interface {
	SetupScraper() *scraper.XScraper
	GetPost(postId string) (*XPost, error)
}

type ScraperService interface {
	Init(username, password string) *scraper.XScraper
	GetPost(postId string) (*XPost, error)
}
