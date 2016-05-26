package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
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

func concurrentQueries() {
	go scrape("http://duckduckgo.com/?q=whitesmith")
	go scrape("http://duckduckgo.com/?q=coimbra")
	go scrape("http://duckduckgo.com/?q=adbjesus")
}

func main() {
	scrape("http://duckduckgo.com/?q=whitesmith")
}
