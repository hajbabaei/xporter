package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"xporter/pkg/scraper"

	"github.com/joho/godotenv"
)

const projectDirName = "bot" // change to relevant project name

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	loadEnv()
	username := os.Getenv("X_USERNAME")
	password := os.Getenv("X_PASSWORD")
	fmt.Println(username)
	// https://github.com/n0madic/twitter-scraper
	fmt.Printf("The username: %s, password: %s\n", username, password)
	scraper := scraper.Init(username, password)
	post := scraper.GetPost("1692588698657865948")
	fmt.Printf("The tweet content: %s\n", post.Text)
}
