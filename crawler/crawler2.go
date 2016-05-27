/*

网页爬虫二：golang抓取iteye首页中的博文推荐栏目里面的文章标题和链接url
http://www.kuaizh.com/?p=639

使用github.com/PuerkitoBio/goquery来爬取网页内容，github.com/PuerkitoBio/goquery的安装详见本人另外一条博文。

Go 实现了类似 jQuery 的功能，包括链式操作语法、操作和查询 HTML 文档。它基于 Go net/html 包和 CSS 选择器库 cascadia。
由于 net/html 解析器返回的是 DOM 节点，而不是完整的 DOM 树，因此，jQuery 的状态操作函数没有实现（像 height()，css()，detach()）。
由于 net/html 解析器要求文档必须是 UTF-8 编码，因此 goquery 库也有此要求。如果文档不是 UTF-8 编码，使用者需要自己转换。进行编码转换，可以使用如下库：
iconv 的 Go 封装，如：github.com/djimenez/iconv-go
官方提供的 text 子仓库，text/encoding，用于其他编码和 UTF-8 之间进行转换

*/

// web_crawler project main.go
// www.kuaizh.com：快智慧，golang网页爬虫
package main

import (
	// "bufio"
	"fmt"
	// "os"

	"github.com/PuerkitoBio/goquery"
)

/*

http://www.iteye.com

              <div class="main_right">
                <h3>
                  <a href="/blogs" target="_blank">博文推荐</a>
                  <a href="/blogs" class="more" target="_blank">[更多]</a>
                </h3>
                <ul>
                                      <li class="title">
                      <a href="http://ye-wolf.iteye.com/blog/2301112" target="_blank" title="青岛JAVA之旅">青岛JAVA之旅</a>
                    </li>
                                      <li class="title">
                      <a href="http://jjhpeopl.iteye.com/blog/2301107" target="_blank" title="mac python2.7安装PIL.Image模块">mac python2.7安装PIL.Image模块</a>
                    </li>

*/

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
	li := ul.Find("li")
	li.Each(func(i int, ss *goquery.Selection) {
		link := *ss.Find("a")
		href, _ := link.Attr("href")
		text := link.Text()
		fmt.Println("text:", text)
		fmt.Println("href:", href)
		fmt.Println()
	})

}

/*

输出内容：

text: 青岛JAVA之旅
href: http://ye-wolf.iteye.com/blog/2301112

text: mac python2.7安装PIL.Image模块
href: http://jjhpeopl.iteye.com/blog/2301107

...

*/
