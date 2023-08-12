package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://info.yorkbbs.ca/list/transfer", g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			// fmt.Println(string(r.Body))
			r.HTMLDoc.Find(".info-item").Each(func(_ int, s *goquery.Selection) {
				fmt.Println(s.Html())
			})
		},
		//BrowserEndpoint: "ws://localhost:3000",
	}).Start()
}
