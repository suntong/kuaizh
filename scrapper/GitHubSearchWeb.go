////////////////////////////////////////////////////////////////////////////
// Package: GitHubSearchWeb.go
// Purpose: GitHub search from its web interface
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

/*

Notes

- The reason for this is that the GitHub code search API needs at least one user, organization, or repository, which defeat the purpose I'm using it
- The GitHub code search does not impose such restrain in its web interface
- BUT! that's from the browswer! -- Tried "https://github.com/search?utf8=%E2%9C%93&q=%22github.com%2Fgoadesign%2Fgoa%2Fdesign%2Fapidsl%22+language%3Ago&type=Code&ref=searchresults" from browser and lynx, or this go code, only web browswer can see the returned results.
- Give up. leave using `scraper.ScrapeHTML` as a reference


Ref

https://godoc.org/gopkg.in/suntong/goscrape.v1

*/

package main

import (
	"os"

	"github.com/go-jsonfile/jsonfile"
	"github.com/go-shaper/shaper"
	"github.com/suntong/goscrape"
	"github.com/suntong/goscrape/extract"
	"github.com/suntong/goscrape/paginate"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var progname = "GitHubSearchWeb"
var buildTime = "2016-12-25"

////////////////////////////////////////////////////////////////////////////
// Function definitions

func main() {
	if len(os.Args) <= 1 {
		println("Usage:\n  ", progname, "github.com_search_string\n\nE.g.\n",
			progname, `'q="github.com/goadesign/goa/design/apidsl"+language:go&type=Code'`)
		os.Exit(0)
	}

	results, err := scrapeWeb("https://github.com/search" + os.Args[1])
	check(err)
	//fmt.Printf("%+v\n", results)

	jsonfile.WriteJSON(os.Stdout, results.Results)
}

func scrapeWeb(url string) (*scrape.ScrapeResults, error) {
	println("Scrapping", url)
	config := &scrape.ScrapeConfig{
		DividePage: scrape.DividePageBySelector("div.code-list > div.code-list-item"),

		Pieces: []scrape.Piece{
			{Name: "FullName", Selector: "p > a:nth-child(1)", Extractor: extract.Text{}},
			{Name: "FileName", Selector: "p > a:nth-child(2)", Extractor: extract.Text{}},
			{Name: "UpdatedAt", Selector: "span.updated-at > relative-time", Extractor: extract.Attr{Attr: "datetime"}},
			{Name: "Language", Selector: "span.language", Extractor: extract.Text{}},
			//{Name: "", Selector: "div.", Extractor: extract.Text{}},
		},

		Paginator:   paginate.BySelector("divp.agination a.next", "href"),
		Opts:        scrape.ScrapeOptions{MaxPages: 1},
		PieceShaper: shaper.NewFilter().ApplyRegSpaces().ApplyTrim(),
	}

	scraper, err := scrape.New(config)
	if err != nil {
		return nil, err
	}

	return scraper.ScrapeHTML(initHTML)
	//return scraper.ScrapeUrl(url)
}

////////////////////////////////////////////////////////////////////////////
// check

func check(e error) {
	if e != nil {
		panic(e)
	}
}

////////////////////////////////////////////////////////////////////////////
// Ref

// https://github.com/search?q="github.com/goadesign/goa/design/apidsl"+language:go&type=Code
// https://github.com/search?q=%22github.com/goadesign/goa/design/apidsl%22+language:go&type=Code
var initHTML = `
<!DOCTYPE html>
<html lang="en" class=" is-copy-enabled is-u2f-enabled">
    <meta charset='utf-8'>
 </head>
  <body class="logged-in  env-production linux">
  <div class="sort-bar">
    <h3>
    We’ve found 257 code results
    </h3>
  </div>

  <div id="code_search_results">
    <div class="code-list">


<div class="code-list-item code-list-item-public ">
    <a href="/james2doyle"><img alt="@james2doyle" class="avatar" height="28" src="https://avatars1.githubusercontent.com/u/1425304?v=3&amp;s=56" width="28"></a>
    <span class="language">
      Go
      
    </span>
  <p class="title">
      <a href="/james2doyle/geo-gambler">james2doyle/geo-gambler</a>
       –
      <a href="/james2doyle/geo-gambler/blob/5dab41b1100d2b7bf04d0c3231c7324c1845b2f9/server/design/api.go" title="server/design/api.go">api.go</a> <br>
      <span class="text-small text-gray match-count">Showing the top six matches.</span>
    <span class="text-small text-gray updated-at">Last indexed <relative-time datetime="2016-10-04T02:13:46Z" title="Oct 3, 2016, 10:13 PM EDT">on Oct 3</relative-time>.</span>
  </p>
    <div class="file-box blob-wrapper">
      <table class="highlight">
          <tbody><tr>
    <td class="blob-num">
      <a href="/james2doyle/geo-gambler/blob/5dab41b1100d2b7bf04d0c3231c7324c1845b2f9/server/design/api.go#L1">1</a>
    </td>
    <td class="blob-code blob-code-inner"><span class="pl-k">package</span> design
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/james2doyle/geo-gambler/blob/5dab41b1100d2b7bf04d0c3231c7324c1845b2f9/server/design/api.go#L2">2</a>
    </td>
    <td class="blob-code blob-code-inner">
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/james2doyle/geo-gambler/blob/5dab41b1100d2b7bf04d0c3231c7324c1845b2f9/server/design/api.go#L3">3</a>
    </td>
    <td class="blob-code blob-code-inner"><span class="pl-k">import</span> (
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/james2doyle/geo-gambler/blob/5dab41b1100d2b7bf04d0c3231c7324c1845b2f9/server/design/api.go#L4">4</a>
    </td>
    <td class="blob-code blob-code-inner">    . <span class="pl-s"><span class="pl-pds">"</span>github.com/goadesign/goa/design<span class="pl-pds">"</span></span>
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/james2doyle/geo-gambler/blob/5dab41b1100d2b7bf04d0c3231c7324c1845b2f9/server/design/api.go#L5">5</a>
    </td>
    <td class="blob-code blob-code-inner">    . <span class="pl-s"><span class="pl-pds">"</span><em>github</em>.<em>com</em>/<em>goadesign</em>/<em>goa</em>/<em>design</em>/<em>apidsl</em><span class="pl-pds">"</span></span></td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/james2doyle/geo-gambler/blob/5dab41b1100d2b7bf04d0c3231c7324c1845b2f9/server/design/api.go#L6">6</a>
    </td>
    <td class="blob-code blob-code-inner">)
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/james2doyle/geo-gambler/blob/5dab41b1100d2b7bf04d0c3231c7324c1845b2f9/server/design/api.go#L7">7</a>
    </td>
    <td class="blob-code blob-code-inner">
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/james2doyle/geo-gambler/blob/5dab41b1100d2b7bf04d0c3231c7324c1845b2f9/server/design/api.go#L8">8</a>
    </td>
    <td class="blob-code blob-code-inner"><span class="pl-k">var</span> <span class="pl-smi">_</span> = <span class="pl-c1">API</span>(<span class="pl-s"><span class="pl-pds">"</span>GeoGambler<span class="pl-pds">"</span></span>, <span class="pl-c1">func</span>() {
</td>
  </tr>

      </tbody></table>
    </div>
</div>


<div class="code-list-item code-list-item-public ">
    <a href="/pablogore"><img alt="@pablogore" class="avatar" height="28" src="https://avatars1.githubusercontent.com/u/3819046?v=3&amp;s=56" width="28"></a>
    <span class="language">
      Go
      
    </span>
  <p class="title">
      <a href="/pablogore/api-test">pablogore/api-test</a>
       –
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/api_definition.go" title="design/api_definition.go">api_definition.go</a> <br>
      <span class="text-small text-gray match-count">Showing the top six matches.</span>
    <span class="text-small text-gray updated-at">Last indexed <relative-time datetime="2016-11-20T23:46:13Z" title="Nov 20, 2016, 6:46 PM EST">on Nov 20</relative-time>.</span>
  </p>
    <div class="file-box blob-wrapper">
      <table class="highlight">
          <tbody><tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/api_definition.go#L1">1</a>
    </td>
    <td class="blob-code blob-code-inner"><span class="pl-k">package</span> design
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/api_definition.go#L2">2</a>
    </td>
    <td class="blob-code blob-code-inner">
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/api_definition.go#L3">3</a>
    </td>
    <td class="blob-code blob-code-inner"><span class="pl-k">import</span> (
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/api_definition.go#L4">4</a>
    </td>
    <td class="blob-code blob-code-inner">	. <span class="pl-s"><span class="pl-pds">"</span>github.com/goadesign/goa/design<span class="pl-pds">"</span></span>
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/api_definition.go#L5">5</a>
    </td>
    <td class="blob-code blob-code-inner">	. <span class="pl-s"><span class="pl-pds">"</span><em>github</em>.<em>com</em>/<em>goadesign</em>/<em>goa</em>/<em>design</em>/<em>apidsl</em><span class="pl-pds">"</span></span></td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/api_definition.go#L6">6</a>
    </td>
    <td class="blob-code blob-code-inner">)
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/api_definition.go#L7">7</a>
    </td>
    <td class="blob-code blob-code-inner">
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/api_definition.go#L8">8</a>
    </td>
    <td class="blob-code blob-code-inner"><span class="pl-k">var</span> <span class="pl-smi">_</span> = <span class="pl-c1">API</span>(<span class="pl-s"><span class="pl-pds">"</span>Gateway<span class="pl-pds">"</span></span>, <span class="pl-c1">func</span>() {
</td>
  </tr>

      </tbody></table>
    </div>
</div>


<div class="code-list-item code-list-item-public ">
    <a href="/pablogore"><img alt="@pablogore" class="avatar" height="28" src="https://avatars1.githubusercontent.com/u/3819046?v=3&amp;s=56" width="28"></a>
    <span class="language">
      Go
      
    </span>
  <p class="title">
      <a href="/pablogore/api-test">pablogore/api-test</a>
       –
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/types_definition.go" title="design/types_definition.go">types_definition.go</a> <br>
      <span class="text-small text-gray match-count">Showing the top six matches.</span>
    <span class="text-small text-gray updated-at">Last indexed <relative-time datetime="2016-11-20T23:46:13Z" title="Nov 20, 2016, 6:46 PM EST">on Nov 20</relative-time>.</span>
  </p>
    <div class="file-box blob-wrapper">
      <table class="highlight">
          <tbody><tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/types_definition.go#L1">1</a>
    </td>
    <td class="blob-code blob-code-inner"><span class="pl-k">package</span> design
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/types_definition.go#L2">2</a>
    </td>
    <td class="blob-code blob-code-inner">
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/types_definition.go#L3">3</a>
    </td>
    <td class="blob-code blob-code-inner"><span class="pl-k">import</span> (
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/types_definition.go#L4">4</a>
    </td>
    <td class="blob-code blob-code-inner">	. <span class="pl-s"><span class="pl-pds">"</span>github.com/goadesign/goa/design<span class="pl-pds">"</span></span>
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/types_definition.go#L5">5</a>
    </td>
    <td class="blob-code blob-code-inner">	. <span class="pl-s"><span class="pl-pds">"</span><em>github</em>.<em>com</em>/<em>goadesign</em>/<em>goa</em>/<em>design</em>/<em>apidsl</em><span class="pl-pds">"</span></span></td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/types_definition.go#L6">6</a>
    </td>
    <td class="blob-code blob-code-inner">)
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/types_definition.go#L7">7</a>
    </td>
    <td class="blob-code blob-code-inner">
</td>
  </tr>
  <tr>
    <td class="blob-num">
      <a href="/pablogore/api-test/blob/fc842d5c960c6e71f43d7a15277d12ce75545fbb/design/types_definition.go#L8">8</a>
    </td>
    <td class="blob-code blob-code-inner"><span class="pl-k">var</span> <span class="pl-smi">MessagePayload</span> = <span class="pl-c1">Type</span>(<span class="pl-s"><span class="pl-pds">"</span>MessagePayload<span class="pl-pds">"</span></span>, <span class="pl-c1">func</span>() {
</td>
  </tr>

      </tbody></table>
    </div>
</div>


</div>

<div class="paginate-container">
        <div class="pagination" data-pjax="true"><span class="previous_page disabled">Previous</span> <em class="current">1</em> <a rel="next" href="/search?p=2&amp;q=%22github.com%2Fgoadesign%2Fgoa%2Fdesign%2Fapidsl%22+language%3Ago&amp;ref=searchresults&amp;type=Code">2</a> <a href="/search?p=3&amp;q=%22github.com%2Fgoadesign%2Fgoa%2Fdesign%2Fapidsl%22+language%3Ago&amp;ref=searchresults&amp;type=Code">3</a> <a href="/search?p=4&amp;q=%22github.com%2Fgoadesign%2Fgoa%2Fdesign%2Fapidsl%22+language%3Ago&amp;ref=searchresults&amp;type=Code">4</a> <a href="/search?p=5&amp;q=%22github.com%2Fgoadesign%2Fgoa%2Fdesign%2Fapidsl%22+language%3Ago&amp;ref=searchresults&amp;type=Code">5</a> <span class="gap">…</span> <a href="/search?p=25&amp;q=%22github.com%2Fgoadesign%2Fgoa%2Fdesign%2Fapidsl%22+language%3Ago&amp;ref=searchresults&amp;type=Code">25</a> <a href="/search?p=26&amp;q=%22github.com%2Fgoadesign%2Fgoa%2Fdesign%2Fapidsl%22+language%3Ago&amp;ref=searchresults&amp;type=Code">26</a> <a class="next_page" rel="next" href="/search?p=2&amp;q=%22github.com%2Fgoadesign%2Fgoa%2Fdesign%2Fapidsl%22+language%3Ago&amp;ref=searchresults&amp;type=Code">Next</a></div>
      </div>

</div>

</body>
</html>
`

/*


 */
