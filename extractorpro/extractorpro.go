package extractorpro

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

type Job struct {
	Title    string
	Company  string
	Ratings  string
	Reviews  string
	URL      string
}

func Extractor() {
	// Open the local HTML file
	file, err := os.Open("page.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse the HTML content
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	// Extract job data
	var jobs []Job
	extractJobs(doc, &jobs)

	// Display extracted job data
	for _, job := range jobs {
		fmt.Println("Title:", job.Title)
		fmt.Println("Company:", job.Company)
		fmt.Println("Ratings:", job.Ratings)
		fmt.Println("Reviews:", job.Reviews)
		fmt.Println("URL:", job.URL)
		fmt.Println()
	}
}

func extractJobs(n *html.Node, jobs *[]Job) {
	if n.Type == html.ElementNode && n.Data == "div" {
		for _, attr := range n.Attr {
			if attr.Key == "class" && strings.Contains(attr.Val, "cust-job-tuple") {
				var job Job

				// Extract data fields
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.Type == html.ElementNode && c.Data == "div" && strings.Contains(getAttributeValue(c, "class"), "row1") {
						for a := c.FirstChild; a != nil; a = a.NextSibling {
							if a.Type == html.ElementNode && a.Data == "a" && strings.Contains(getAttributeValue(a, "class"), "title") {
								job.Title = strings.TrimSpace(a.FirstChild.Data)
								job.URL = getAttributeValue(a, "href")
							}
						}
					} else if c.Type == html.ElementNode && c.Data == "div" && strings.Contains(getAttributeValue(c, "class"), "row2") {
						for span := c.FirstChild; span != nil; span = span.NextSibling {
							if span.Type == html.ElementNode && span.Data == "span" && strings.Contains(getAttributeValue(span, "class"), "comp-dtls-wrap") {
								for a := span.FirstChild; a != nil; a = a.NextSibling {
									if a.Type == html.ElementNode && a.Data == "a" && strings.Contains(getAttributeValue(a, "class"), "comp-name") {
										job.Company = strings.TrimSpace(a.FirstChild.Data)
									}
								}
							} else if span.Type == html.ElementNode && span.Data == "a" && strings.Contains(getAttributeValue(span, "class"), "rating") {
								job.Ratings = strings.TrimSpace(span.FirstChild.NextSibling.Data)
							} else if span.Type == html.ElementNode && span.Data == "a" && strings.Contains(getAttributeValue(span, "class"), "review") {
								job.Reviews = strings.TrimSpace(span.FirstChild.Data)
							}
						}
					}
				}

				// Append job to the list
				*jobs = append(*jobs, job)
			}
		}
	}

	// Recursively traverse the HTML tree
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractJobs(c, jobs)
	}
}

func getAttributeValue(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}
