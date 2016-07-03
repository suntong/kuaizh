////////////////////////////////////////////////////////////////////////////
// Program: kuaizh1.go
// Purpose: kuaizh web scrapping tool
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-shaper/shaper"
	"github.com/spakin/awk"
	"github.com/suntong/goscrape"
	"github.com/suntong/goscrape/extract"
	"github.com/suntong/goscrape/paginate"
	"gopkg.in/pipe.v2"
)

func main() {
	results, err := scrapeIndexes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scraping: %s\n", err)
		os.Exit(1)
	}

	items := []string{}
	for _, rs := range results.Results {
		for i, r := range rs {
			fmt.Printf("%d: %+v\n", i, r)
			items = append(items, r["Link"].(string))
		}
	}

	fmt.Println("\n")
	// cut out the last 7 items
	items = items[:len(items)-7]
	// duplicate the last item
	items = append(items, items[len(items)-1])
	scrapePages(items)
}

func scrapeIndexes() (*scrape.ScrapeResults, error) {
	config := &scrape.ScrapeConfig{
		DividePage: scrape.DividePageBySelector("article.post"),

		Pieces: []scrape.Piece{
			{Name: "Title", Selector: "h2.entry-title", Extractor: extract.Text{}},
			{Name: "Link", Selector: "h2.entry-title a", Extractor: extract.Attr{Attr: "href"}},
			{Name: "Description", Selector: "div.entry-summary p", Extractor: extract.Text{}},
			//{Name: "", Selector: "div.", Extractor: extract.Text{}},
		},

		Paginator: paginate.BySelector("div#pagination a.next", "href"),
		//Opts:        scrape.ScrapeOptions{MaxPages: 1},
		PieceShaper: shaper.NewFilter().ApplyRegSpaces(), // .ApplyTrim()
	}

	scraper, err := scrape.New(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating scraper: %s\n", err)
		os.Exit(1)
	}

	//return scraper.ScrapeHTML(initHTMLI)
	return scraper.ScrapeUrl("http://www.kuaizh.com/?cat=12")
}

// Input is url slice
func scrapePages(items []string) {
	ilen := len(items)
	fmt.Println("Total: ", ilen)
	for i := range items {
		// iterate over the slice in reverse order
		ir := i + 1
		url := items[ilen-ir]
		if url != "" {
			id := fmt.Sprintf("%02d", ir)
			title := scrapePage(id, url)
			fmt.Println("TT:", title)
			time.Sleep(2 * time.Second)
		}
	}
}

func scrapePage(id, url string) (title string) {
	fmt.Printf("\n%s: %s\n", id, url)
	// func NewDocumentFromReader(r io.Reader) (*Document, error)
	//doc, _ := goquery.NewDocumentFromReader(strings.NewReader(initHTMLP))
	doc, _ := goquery.NewDocument(url)
	title = doc.Find("article.post h1.entry-title").Text()
	cnt := doc.Find("div.entry-content")
	cntStr := cnt.Text()

	r := strings.NewReplacer(":", "：", "/", "／", " ", "")
	outfile := fmt.Sprintf("%s-%s.go", id, r.Replace(title))
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s\n%s\n\n", title, url)
	buf.WriteString(cntStr)
	psrc, pdst := pipe.Script(
		pipe.Read(buf),
	), pipe.Script(
		pipe.WriteFile(outfile, 0644),
	)

	p := pipe.Line(
		psrc,
		pgAwk(),
		pdst,
	)
	err := pipe.Run(p)
	check(err)
	return
}

func pgAwk() pipe.Pipe {
	return pipe.TaskFunc(func(st *pipe.State) error {
		// == Setup
		s := awk.NewScript()
		s.Output = st.Stdout

		s.Begin = func(s *awk.Script) {
			s.Println("/*\n")
		}

		s.End = func(s *awk.Script) {
			s.Println("\n*/")
		}

		s.AppendStmt(func(s *awk.Script) bool {
			return s.F(1).Match("^ *package")
		}, func(s *awk.Script) {
			s.Println("\n*/\n")
		})

		// s.AppendStmt(func(s *awk.Script) bool {
		// 	return s.F(1).Match("^}")
		// }, func(s *awk.Script) {
		// 	s.Println()
		// 	s.Println("\n/*\n")
		// 	s.Next()
		// })
		s.AppendStmt(func(s *awk.Script) bool {
			return s.F(1).Match("^ *(执行|输出|结果)")
		}, func(s *awk.Script) {
			s.Println("/*\n")
		})

		// 1; # i.e., print all
		s.AppendStmt(nil, nil)

		// == Run it
		return s.Run(st.Stdin)
	})

}

//::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
// Support functions

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// http://www.kuaizh.com/?cat=12
var initHTMLI = `
`

// http://www.kuaizh.com/?p=639
var initHTMLP = `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width, initial-scale=1.0"> 		
    <title>golang第五十八天:go语言中的程序退出Exit | 快智慧程序员社区</title>
</head>

<body class="single single-post postid-781 single-format-standard group-blog full-width singular">

<article id="post-781" class="post-781 post type-post status-publish format-standard hentry category-62day-learning-golang">
	<header class="entry-header">
		<div class="entry-thumbnail">
		    <a href="http://www.kuaizh.com/?p=781">
			   		    </a>
		</div>
		<h1 class="entry-title">golang第五十八天:go语言中的程序退出Exit</h1>
	</header><!-- .entry-header -->

	<div class="entry-content">
		<pre>

使用os.Exit来退出程序，返回状态码。

package main


import "fmt"
import "os"

	

func main() {

	//当程序退出时，defer语句不会被执行。
    defer fmt.Println("!")
	//退出程序，返回状态码3
    os.Exit(3)
}

执行结果：
$ go run exit.go
exit status 3
</pre>
			</div><!-- .entry-content -->

	<footer class="entry-meta">
		This entry was posted in <a href="http://www.kuaizh.com/?cat=12" rel="category">62天golang学习笔记</a>. Bookmark the <a href="http://www.kuaizh.com/?p=781" title="Permalink to golang第五十八天:go语言中的程序退出Exit" rel="bookmark">permalink</a>.
			</footer><!-- .entry-meta -->
</article><!-- #post-## -->

</body>
</html>
`

/*


 */
