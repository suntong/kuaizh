package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, _ := goquery.NewDocument("http://syui.github.io/")
	doc.Find(".post h1 a").Each(func(_ int, s *goquery.Selection) {
		if articleUrl, ok := s.Attr("href"); ok {
			fmt.Println(articleUrl)
		}
	})
}
