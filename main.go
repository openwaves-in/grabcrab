package main

import "github.com/openwaves-in/grabcrab/urlcooker"

//"github.com/openwaves-in/grabcrab/crawlerpro"
//"github.com/openwaves-in/grabcrab/extractorpro"

func main() {
	//	url := "https://www.naukri.com/java-fresher-jobs-in-chennai?k=java%20fresher&l=chennai&experience=0"

	//crawlerpro.Crawl("https://www.naukri.com/java-fresher-jobs-in-chennai-3?experience=0")
	//extractorpro.Extractor()
	urlcooker.Urlcook("java-fresher", "chennai")
}
