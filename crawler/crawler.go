package crawler

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.9999.999 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299",
	// Add more user agents as needed
}

func getRandomUserAgent() string {
	rand.Seed(time.Now().UnixNano())
	return userAgents[rand.Intn(len(userAgents))]
}

func Crawl(url string) {
	// URL to crawl
	

	// Create a new collector
	c := colly.NewCollector()

	// Randomize user agent for each request
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", getRandomUserAgent())
	})

	// On HTML response, save HTML content to a file
	c.OnResponse(func(r *colly.Response) {
		err := os.WriteFile("page.html", r.Body, 0644)
		if err != nil {
			fmt.Println("Error saving HTML content to file:", err)
			return
		}
		fmt.Println("Page saved successfully as page.html")
	})

	// Start scraping
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error visiting URL:", err)
		return
	}
}
