package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func scrape(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a.result__url").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if !exists {
			return
		}
		fmt.Println(link)
	})
}

func main() {
	scrape("http://duckduckgo.com/?q=whitesmith")
}
