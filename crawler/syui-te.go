package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	var (
		stringFlag string
	)
	flag.StringVar(&stringFlag, "string", "blank", "string flag")
	flag.StringVar(&stringFlag, "s", "blank", "string flag")
	flag.Parse()

	doc, _ := goquery.NewDocument(stringFlag)
	doc.Find(".post h1 a").Each(func(_ int, s *goquery.Selection) {
		if articleUrl, ok := s.Attr("href"); ok {
			fmt.Println(articleUrl)
		}
	})
}
