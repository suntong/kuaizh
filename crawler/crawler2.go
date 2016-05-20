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
	"bufio"
	"fmt"
	"os"

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
	li := ul.Find("li")
	li.Each(func(i int, ss *goquery.Selection) {
		link := *ss.Find("a")
		href, _ := link.Attr("href")
		text := link.Text()
		fmt.Println("text:", text)
		fmt.Println("href:", href)
	})

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		fmt.Println(url)
	}

}

/*

输出内容：
C:/Go/bin/go.exe build -i [C:/Users/Administrator/web_crawler/src/web_crawler]
成功: 进程退出代码 0.
C:/Users/Administrator/web_crawler/src/web_crawler/web_crawler.exe  [C:/Users/Administrator/web_crawler/src/web_crawler]
text: 026_CoreAPI_Configuration_SessionFactory ...
href: http://yuzhouxiner.iteye.com/blog/2268457
text: RedHatLinux6.5下安装无线网卡驱动
href: http://wjrko.iteye.com/blog/2268451
text: Linux Shell
href: http://gengzg.iteye.com/blog/2268448
text:  自己项目中PHP常用工具类大全分享
href: http://shouce.iteye.com/blog/2268404
text: 自己实现动态代理
href: http://hangzhoujava.iteye.com/blog/2268400
text: 探究数值比较在程序中执行效率的区别
href: http://tangl163.iteye.com/blog/2268327
text: Michael Nielsen 's 神经网络学习之一
href: http://luchi007.iteye.com/blog/2268309
text: 020_ID生成策略_XML_配置
href: http://yuzhouxiner.iteye.com/blog/2268315
text: 【MongoDB】的安装与基本操作
href: http://gaojingsong.iteye.com/blog/2268304
text: 常用git命令总结
href: http://itxiaojiang.iteye.com/blog/2268253
text: 系统性能优化方法
href: http://itxiaojiang.iteye.com/blog/2268252
text: OkHttp的使用简介及封装，实现更简洁的调用 ...
href: http://dzc.iteye.com/blog/2268386
text: 归来的微软
href: http://jiezhu2007.iteye.com/blog/2268117
text: JSP的几种参数传值
href: http://gaojingsong.iteye.com/blog/2268297
text: JAVA代码操作Memcache
href: http://gaojingsong.iteye.com/blog/2268164

*/
