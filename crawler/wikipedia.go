package main

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Stock struct {
	Symbol string
	CIK    int64
}

func main() {
	stocks := make([]Stock, 0, 600)
	doc, _ := goquery.NewDocument("https://en.wikipedia.org/wiki/List_of_S%26P_500_companies")
	doc.Find("table.wikitable").First().Find("tr").Each(func(i int, s *goquery.Selection) {
		cik, _ := strconv.ParseInt(s.Find("td").Last().Text(), 10, 64)
		stocks = append(stocks, Stock{
			Symbol: s.Find("td a").First().Text(),
			CIK:    cik,
		})
	})
	for _, stock := range stocks {
		fmt.Println(stock)
	}

}

/*

<table class="wikitable sortable">

<tr>
<th><a href="/wiki/Ticker_symbol" title="Ticker symbol">Ticker symbol</a></th>
<th>Security</th>
<th><a href="/wiki/SEC_filing" title="SEC filing">SEC filings</a></th>
<th><a href="/wiki/Global_Industry_Classification_Standard" title="Global Industry Classification Standard">GICS</a> Sector</th>
<th>GICS Sub Industry</th>
<th>Address of Headquarters</th>
<th>Date first added</th>
<th><a href="/wiki/Central_Index_Key" title="Central Index Key">CIK</a></th>
</tr>

<tr>
<td><a rel="nofollow" class="external text" href="https://www.nyse.com/quote/XNYS:MMM">MMM</a></td>
<td><a href="/wiki/3M" title="3M">3M Company</a></td>
<td><a rel="nofollow" class="external text" href="http://www.sec.gov/cgi-bin/browse-edgar?CIK=MMM&amp;action=getcompany">reports</a></td>
<td>Industrials</td>
<td>Industrial Conglomerates</td>
<td><a href="/wiki/St._Paul,_Minnesota" title="St. Paul, Minnesota" class="mw-redirect">St. Paul, Minnesota</a></td>
<td></td>
<td>0000066740</td>
</tr>

*/
