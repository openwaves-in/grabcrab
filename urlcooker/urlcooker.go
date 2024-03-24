package urlcooker

// url style https://www.naukri.com/golang-developer-jobs-in-hydeabad?k=golang%20developer&l=hydeabad&experience=0
// page 2    https://www.naukri.com/golang-developer-jobs-in-hydeabad-2?k=golang+developer&l=hydeabad&experience=0

import (
	"fmt"

	"github.com/openwaves-in/grabcrab/crawlerpro"
	"github.com/openwaves-in/grabcrab/extractorpro"
	"github.com/openwaves-in/grabcrab/pathfinder"
)

func generateLinks(baseLink string, count int) []string {
	links := make([]string, count)
	for i := 0; i < count; i++ {
		if i == 0 {
			links[i] = baseLink
		} else {
			links[i] = fmt.Sprintf("%s-%d", baseLink, i)
		}
	}
	return links
}

//java-fresher-jobs-in-chennai?k=java%20fresher&l=chennai&experience=0

func Urlcook(skilss string, location string) string {
	baseLink := "https://www.naukri.com/" + skilss + "-jobs-in-" + location
	addLink := "?jobAge=7&experience=0"
	links := generateLinks(baseLink, 5)
	for _, link := range links {
		url := link + addLink
		fmt.Println(url)
		crawlerpro.Crawl(url)
		extractorpro.Extractor()

	}
	return pathfinder.Fpath() 
}
