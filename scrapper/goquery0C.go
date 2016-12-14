// https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks

package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("http://www.factorydirect.ca/Coupons.aspx")
	if err != nil {
		log.Fatal(err)
	}

	sel := doc.Find("body")
	for i := range sel.Nodes {
		s := sel.Eq(i)
		// use `s` as a single selection of 1 node
		f, _ := s.Find("form").Attr("action") // :first-child
		fmt.Printf("%d) %+v\n", i+1, f)
	}
}
