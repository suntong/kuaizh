package main

import (
	"fmt"
	"os"

	"github.com/andrew-d/goscrape"
	"github.com/andrew-d/goscrape/extract"
	"github.com/andrew-d/goscrape/paginate"
)

func main() {
	config := &scrape.ScrapeConfig{
		DividePage: scrape.DividePageBySelector("tr.athing"),

		Pieces: []scrape.Piece{
			{Name: "title", Selector: "td.title > a", Extractor: extract.Text{}},
			{Name: "link", Selector: "td.title > a", Extractor: extract.Attr{Attr: "href"}},
			{Name: "no", Selector: "span.rank", Extractor: extract.Text{}},
			// Extractor: extract.Regex{Regex: regexp.MustCompile(`(\d+)`)}}
		},

		Paginator: paginate.BySelector("a.morelink", "href"),
	}

	scraper, err := scrape.New(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating scraper: %s\n", err)
		os.Exit(1)
	}

	results, err := scraper.ScrapeWithOpts(
		"https://news.ycombinator.com",
		scrape.ScrapeOptions{MaxPages: 3},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scraping: %s\n", err)
		os.Exit(1)
	}

	//json.NewEncoder(os.Stdout).Encode(results)
	//for _, result := range results.Results {
	//	for _, r := range result {
	// for i, r := range results.Results[0] {
	// 	//json.NewEncoder(os.Stdout).Encode(r)
	// 	fmt.Printf("%d %s %s\n%s\n\n", i, r["no"], r["title"], r["link"])
	// }
	//}
	for _, r := range results.Results[0] {
		fmt.Printf("%s %s\n%s\n\n", r["no"], r["title"], r["link"])
	}
}

/*

1. Alan Kay's reading list for his students
http://www.squeakland.org/resources/books/readingList.jsp

...

28. How to print things
https://byorgey.wordpress.com/how-to-print-things/

29. A Ruby wrapper for LaTeXML
https://github.com/Authorea/latexml-ruby

30. GE's Walking Truck (1969) [video]
http://www.educatedearth.net/video.php?id=5000

*/
