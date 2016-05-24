// web_crawler project main.go
// www.kuaizh.com：快智慧，golang网页爬虫
// http://www.kuaizh.com/?p=641

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	//获取html的document对象
	doc, _ := goquery.NewDocument("http://www.iteye.com")
	//fmt.Println(doc.Html())
	//根据样式查询博文推荐栏目
	mainRight := doc.Find(".main_right")
	//fmt.Println(mainRight.Html())
	ul := mainRight.Find("ul")
	//fmt.Println(ul.Html())

	//解析每一条博文的标题和url链接
	items := make([]string, 2)
	li := ul.Find("li")
	li.Each(func(i int, ss *goquery.Selection) {
		link := *ss.Find("a")
		href, _ := link.Attr("href")
		text := link.Text()
		fmt.Println("text:", text)
		fmt.Println("href:", href)
		items = append(items, href)
	})

	crawler(items)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	url := scanner.Text()
	fmt.Println(url)

}

//参数为要爬取的url分片
func crawler(items []string) {
	fmt.Println("爬取博文内容条数：", len(items))
	time.Sleep(3 * time.Second)
	for _, url := range items {
		fmt.Println("开始抓取url：", url)
		if url != "" {
			doc, _ := goquery.NewDocument(url)
			title := doc.Find(".blog_title").Find("h3").Find("a").Text()
			fmt.Println("标题:", title)
			cnt := doc.Find(".blog_content")
			cntStr := cnt.Text()
			fmt.Println("内容:", cntStr)
			time.Sleep(2 * time.Second)
		}
	}
}
