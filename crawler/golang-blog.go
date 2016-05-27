package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

/*

https://blog.golang.org

	<div class="article">
		<h3 class="title"><a href="/go1.6">Go 1.6 is released</a></h3>
		<p class="date">17 February 2016</p>

*/

func main() {
	doc, err := goquery.NewDocument("https://blog.golang.org")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".article").Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3").Text()
		link, _ := s.Find("h3 a").Attr("href")
		fmt.Printf("%d) %s - %s\n", i+1, title, link)
	})
}

/*

1) Go 1.6 is released - /go1.6
2) Language and Locale Matching in Go - /matchlang
3) Six years of Go - /6years
4) GolangUK 2015 - /gouk15
5) Go GC: Prioritizing low latency and simplicity - /go15gc

*/
