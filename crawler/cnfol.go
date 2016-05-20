package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	g, e := goquery.NewDocument("http://gold.3g.cnfol.com/")
	if e != nil {
		fmt.Println(e)
	}
	c := g.Find("ul")
	s := c.Eq(6).Find("a")
	s.Each(func(i int, content *goquery.Selection) {
		a, _ := content.Attr("href")
		fmt.Println(a)
	})
}
