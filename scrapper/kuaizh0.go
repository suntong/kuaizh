////////////////////////////////////////////////////////////////////////////
// Program: scrape_kijiji_car-r
// Purpose: Kijiji web scrapping using the Raw interface
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/go-shaper/shaper"
	"github.com/suntong/goscrape"
	"github.com/suntong/goscrape/extract"
	"github.com/suntong/goscrape/paginate"
)

func main() {
	results, err := scrapeIndexes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scraping: %s\n", err)
		os.Exit(1)
	}

	for _, rs := range results.Results {
		for i, r := range rs {
			fmt.Printf("%d: %+v\n", i, r)
		}
	}
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

		Paginator:   paginate.BySelector("div#pagination a.next", "href"),
		Opts:        scrape.ScrapeOptions{MaxPages: 1},
		PieceShaper: shaper.NewFilter().ApplyRegSpaces(), // .ApplyTrim()
	}

	scraper, err := scrape.New(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating scraper: %s\n", err)
		os.Exit(1)
	}

	return scraper.ScrapeHTML(initHTMLI)
}

// http://www.kuaizh.com/?cat=12
var initHTMLI = `
<html lang="zh-CN">
<head>
	<meta charset="UTF-8">
    <title>62天Golang学习笔记 | 快智慧程序员社区</title>
  </head>

<body class="archive category category-62day-learning-golang category-12 group-blog full-width">

<div class="container">
	<div class="row" role="main">
    <div class="span8">

		
			<header class="page-header">
				<h1 class="page-title">
					62天golang学习笔记				</h1>
							</header><!-- .page-header -->

						
				
<article id="post-781" class="post-781 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=781" rel="bookmark">golang第五十八天:go语言中的程序退出Exit</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=781">
		   		</a>	
	</div>
		<p>使用os.Exit来退出程序，返回状态码。 package main import “fmt” import “os” func main() { //当程序退出时，de…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=781">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-779" class="post-779 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=779" rel="bookmark">golang第五十七天:go语言中的处理系统信号量Signals</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=779">
		   		</a>	
	</div>
		<p>有时候我们需要处理系统的信号量，在实际项目中我们可能有下面的需求： 1、修改了配置文件后，希望在不重启进程的情况下重新加载配置文件； 2、当用 Ctrl + C 强制关闭…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=779">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-777" class="post-777 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=777" rel="bookmark">golang第五十六天:go语言中的执行进程Exec’ing Processes</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=777">
		   		</a>	
	</div>
		<p>有时候，我们只想用其他的（也许是非 Go 程序）来完全替代当前的 Go 进程。这时候，我们可以使用经典的 exec方法的 Go 实现。 package main impo…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=777">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-775" class="post-775 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=775" rel="bookmark">golang第五十五天:go语言中的启动进程以及调用shell等系统命令Spawning Processes</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=775">
		   		</a>	
	</div>
		<p>有时候我们的go程序需要调用一些外部进程，比如系统命令，shell脚本等，使用exec包可以创建一个外部进程。 package main import “fmt” imp…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=775">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-773" class="post-773 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=773" rel="bookmark">golang第五十四天:go语言中的获取和设置环境变量Environment Variables</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=773">
		   		</a>	
	</div>
		<p>Golang 要获取环境变量需要使用os包。导入”os”包，通过os包中的Getenv方法来获取。Setenv设置环境变量 package main import “os…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=773">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-771" class="post-771 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=771" rel="bookmark">golang第五十三天:go语言中的命令行标志Command-Line Flags</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=771">
		   		</a>	
	</div>
		<p>命令行标志（使用标记的命令行处理方法，命令行标志是命令行程序指定选项的常用方式。例如，在 wc -l 中，这个 -l 就是一个命令行标志。） package main /…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=771">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-769" class="post-769 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=769" rel="bookmark">golang第五十二天:go语言中的命令行参数Command-Line Arguments</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=769">
		   		</a>	
	</div>
		<p>命令行参数(命令行参数是指定程序运行参数的一个常见方式。例如，go run hello.go，程序 go 使用了 run 和 hello.go 两个参数。) packag…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=769">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-767" class="post-767 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=767" rel="bookmark">golang第五十一天:go语言中的Line Filters</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=767">
		   		</a>	
	</div>
		<p>行过滤是一种常用的编程方式，它读取标准输入，进行转换，然后把结果打印到标准输出。比如linux的grep和sed。 package main import ( “bufi…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=767">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-765" class="post-765 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=765" rel="bookmark">golang第五十一天:go语言中的文件操作之写入文件Writing Files</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=765">
		   		</a>	
	</div>
		<p>Golang简单写文件操作的四种方法： 第一种方式:使用io.WriteString写入文件 第二种方式:使用ioutil.WriteFile 第三种方式:使用File(…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=765">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-763" class="post-763 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=763" rel="bookmark">golang第五十天:go语言中的文件操作之读取文件Reading Files</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=763">
		   		</a>	
	</div>
		<p>读写文件是最基本的功能。go语言读文件挺有意思，由于go语言的interface，使得go语言与其他语言有所不同。 与其他语言一样，go语言有File类型的结构体，但Fi…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=763">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-761" class="post-761 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=761" rel="bookmark">golang第四十九天:go语言中的Base64编码Base64 Encoding</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=761">
		   		</a>	
	</div>
		<p>go语言对Base64编码提供了内置的支持。 package main //这个语法引入了 encoding/base64 包并使用名称 b64代替默认的 base64。…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=761">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-759" class="post-759 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=759" rel="bookmark">golang第四十八天:go语言中的sha1加密SHA1 Hashes</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=759">
		   		</a>	
	</div>
		<p>SHA-1是一种数据加密算法，该算法的思想是接收一段明文，然后以一种不可逆的方式将它转换成一段（通常更小）密文， 也可以简单的理解为取一串输入码（称为预映射或信息），并把…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=759">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-757" class="post-757 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=757" rel="bookmark">golang第四十八天:go语言中的字符串转换成URL对象URL Parsing</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=757">
		   		</a>	
	</div>
		<p>URL提供了一种统一访问资源的方式。我们来看一下Go里面如何解析URL。 url包解析URL并实现了查询的逸码，参见RFC 3986。 func Parse(rawurl…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=757">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-755" class="post-755 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=755" rel="bookmark">golang第四十七天:go语言中的字符串转换成数字Number Parsing</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=755">
		   		</a>	
	</div>
		<p>把字符串转换成数字是最常用的功能。go语言的strconv包提供了把字符串转换成数字的功能。 package main import “strconv” import “…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=755">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-753" class="post-753 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=753" rel="bookmark">golang第四十六天:go语言中的随机数生成方法Random Numbers</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=753">
		   		</a>	
	</div>
		<p>golang生成随机数可以使用math/rand包 package main import “time” import “fmt” import “math/rand” …</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=753">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-751" class="post-751 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=751" rel="bookmark">golang第四十五天:go语言中的日期时间对象和字符串对象之间的转换Time Formatting / Parsing</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=751">
		   		</a>	
	</div>
		<p>Go语言中，获取时间戳用time.Now().Unix()，获取时间对象用time.Now()，格式化时间用t.Format，解析时间用time.Parse。 在Go语言…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=751">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-749" class="post-749 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=749" rel="bookmark">golang第四十四天:go语言中的unix时间戳Epoch</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=749">
		   		</a>	
	</div>
		<p>通常我们需要获取unix时间戳，即从1970年1月1日（UTC/GMT的午夜）开始所经过的秒数，毫秒，纳秒等。 一个小时表示为UNIX时间戳格式为：3600秒；一天表示为…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=749">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-747" class="post-747 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=747" rel="bookmark">golang第四十三天:go语言中的日期时间操作Time</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=747">
		   		</a>	
	</div>
		<p>time处理包括俩个方面：时间点（某一时刻）和时长（某一段时间） 编程离不开时间，时间管理，严格的说分成两块，一个是当前的时刻，对应的是一个点，还有是一段时间间隔。 pa…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=747">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-745" class="post-745 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=745" rel="bookmark">golang第四十二天:go语言中的JSON数据操作</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=745">
		   		</a>	
	</div>
		<p>Go语言转换JSON数据真是非常的简单。Go按照RFC 4627的标准实现了一个json编解码的标准库。 Marshal 用于将对象序列化到json对象中，Unmarsh…</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=745">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
				
<article id="post-743" class="post-743 post type-post status-publish format-standard hentry category-62day-learning-golang">
	
	<header class="entry-header">
		<h2 class="entry-title"><a href="http://www.kuaizh.com/?p=743" rel="bookmark">golang第四十一天:go语言中的正则表达式Regular Expressions</a></h2>	</header><!-- .entry-header -->
	<div class="entry-summary">
	<div class="summary-thumbnail">
		<a href="http://www.kuaizh.com/?p=743">
		   		</a>	
	</div>
		<p>Golang中的正则表达式 对于 [a-z] 这样的正则表达式，如果要在 [] 中匹配 – ，可以将 – 放在 [] 的开头或结尾，例如 [-a-z] 或 [a-z-] …</p>
				<!--
		<div class="read-more button">
		    <a href="http://www.kuaizh.com/?p=743">阅读全文 &raquo;</a>		</div>
		-->
		
	</div><!-- .entry-summary -->
	<div class="clearfix"></div>    
				
</article><!-- #post-## -->
			
			<div id="pagination"><div class="total-pages btn special">Page 1 of 4</div><div class="btn-group"><span class="page-numbers btn disabled">1</span>
<a class="page-numbers btn" href="http://www.kuaizh.com/?cat=12&amp;paged=2">2</a>
<a class="page-numbers btn" href="http://www.kuaizh.com/?cat=12&amp;paged=3">3</a>
<a class="page-numbers btn" href="http://www.kuaizh.com/?cat=12&amp;paged=4">4</a>
<a class="next page-numbers btn" href="http://www.kuaizh.com/?cat=12&amp;paged=2">»</a></div></div>
		
		</div>
  </div>
</div><!-- #content -->
</body>
</html>
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
