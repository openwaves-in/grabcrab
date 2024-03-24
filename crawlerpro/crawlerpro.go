package crawlerpro

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/chromedp/chromedp"
)

func Crawl(siteurl string) {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Create a timeout to prevent hanging
	ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	// Create options
	options := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),      // Set headless mode to false for better visibility
		chromedp.UserAgent(RandomUserAgent()), // Randomize user agent
	)

	// Create a new context with the created options
	ctx, cancel = chromedp.NewExecAllocator(ctx, options...)
	defer cancel()

	// Create a new browser context
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	// Navigate to the URL
	if err := chromedp.Run(ctx,
		chromedp.Navigate(siteurl),
		chromedp.Sleep(time.Duration(rand.Intn(5)+3)*time.Second), // Simulate human-like behavior
	); err != nil {
		log.Fatal(err)
	}

	// Get the HTML content
	var htmlContent string
	if err := chromedp.Run(ctx, chromedp.OuterHTML("html", &htmlContent)); err != nil {
		log.Fatal(err)
	}

	// Save the HTML content to a file
	if err := ioutil.WriteFile("page.html", []byte(htmlContent), 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Page saved successfully")

}

// RandomUserAgent generates a random user-agent string
func RandomUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:97.0) Gecko/20100101 Firefox/97.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Safari/537.36",
	}
	return userAgents[rand.Intn(len(userAgents))]
}
