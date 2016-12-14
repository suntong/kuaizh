// https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks

package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://blog.golang.org")
	if err != nil {
		log.Fatal(err)
	}

	sel := doc.Find(".article")
	for i := range sel.Nodes {
		s := sel.Eq(i)
		// use `s` as a single selection of 1 node
		title := s.Find("h3").Text()
		link, _ := s.Find("h3 a").Attr("href")
		fmt.Printf("%d) %s - %s\n", i+1, title, link)
	}
}

/*

<div id="content">
  <h1><a href="/">The Go Blog</a></h1>
  <div class="article">
    <h3 class="title">
      <a href="/survey2016">
        Participate in the 2016 Go User Survey and Company Questionnaire</a>
      </h3>
    <p class="date">13 December 2016</p>
...

*/
