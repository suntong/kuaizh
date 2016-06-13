package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-shaper/shaper"
	"github.com/suntong/goscrape"
	"github.com/suntong/goscrape/extract"
	"github.com/suntong/goscrape/paginate"
)

func main() {
	config := &scrape.ScrapeConfig{
		DividePage: scrape.DividePageBySelector("div.info-container"),

		Pieces: []scrape.Piece{
			{Name: "FullHtml", Selector: ".", Extractor: extract.OuterHtml{}},
			{Name: "Price", Selector: "div.price", Extractor: extract.Text{}},
			// {Name: "Id", Selector: "div.title>a.attr('href')",
			// 	Extractor: extract.Regex{Regex: regexp.MustCompile(`.*/(\d+)$`)}},
			{Name: "Title", Selector: "div.title", Extractor: extract.Text{}},
			{Name: "Link", Selector: "div.title > a", Extractor: extract.Attr{Attr: "href"}},
			{Name: "Description", Selector: "div.description", Extractor: extract.Text{}},
			{Name: "Details", Selector: "div.details", Extractor: extract.Text{}},
			//{Name: "", Selector: "div.", Extractor: extract.Text{}},
		},

		Paginator:   paginate.BySelector("a[title='Next']", "href"),
		PieceShaper: shaper.NewFilter().ApplyTrim().ApplyRegSpaces(),
	}

	scraper, err := scrape.New(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating scraper: %s\n", err)
		os.Exit(1)
	}

	results, err := scraper.ScrapeWithOpts(
		"http://www.kijiji.ca/b-cars-trucks/gta-greater-toronto-area/gps/k0c174l1700272?ad=offering&price=__18000&kilometers=__90000&transmission=2&for-sale-by=ownr",
		scrape.ScrapeOptions{MaxPages: 5},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scraping: %s\n", err)
		os.Exit(1)
	}

	json.NewEncoder(os.Stdout).Encode(results)
}
