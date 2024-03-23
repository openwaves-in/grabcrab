package main

import (
	"github.com/openwaves-in/grabcrab/crawler"
	"github.com/openwaves-in/grabcrab/extractorpro"
)

func main() {
	url := "https://www.naukri.com/java-fresher-jobs-in-chennai?k=java%20fresher&l=chennai&experience=0"

	crawler.Crawl(url)
	extractorpro.Extractor()
}
