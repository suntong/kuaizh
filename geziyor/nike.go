package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://www.nike.com/jp/launch?s=in-stock", g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			log.Println(string(r.Body))
			r.HTMLDoc.Find("figure").Each(func(_ int, s *goquery.Selection) {
				log.Println(s.Text())
			})
		},
		//BrowserEndpoint: "ws://localhost:3000",
	}).Start()
}
