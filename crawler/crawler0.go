package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("http://sports.sina.com.cn")
	//fmt.Println(doc.Html())
	if err != nil {
		panic(err)
	} else {
		pTitle := doc.Find("title").Text() //直接提取title的内容
		fmt.Println(pTitle)
		nav := doc.Find(".sina15-nav-list")
		fmt.Println(nav.Html())
	}
}

/*

新浪体育_新浪网

        <li class="sport-logo"><a href="http://sports.sina.com.cn/"><img alt="新浪体育" src="http://n.sinaimg.cn/sports/index1508/logo.png" width="120" height="30"/></a></li>
        <li class="sina15-nav-list-first"><a href="http://www.sina.com.cn/">新浪首页</a></li>
        <li><a href="http://news.sina.com.cn/">新闻</a></li>
        <li><a href="http://sports.sina.com.cn/">体育</a></li>
        <li><a href="http://finance.sina.com.cn/">财经</a></li>
        <li><a href="http://ent.sina.com.cn/">娱乐</a></li>
        <li><a href="http://tech.sina.com.cn/">科技</a></li>
        <li><a href="http://blog.sina.com.cn/">博客</a></li>
        <li><a href="http://photo.sina.com.cn/">图片</a></li>
        <li><a href="http://zhuanlan.sina.com.cn/">专栏</a></li>
        <li class="sina15-nav-list-last"><a href="#" class="sina16-more" data-action="dropdown" data-target="more"><i class="sina15-icon sina15-icon-arrows-a sina15-icon-arrows-b"></i></a></li>
       <nil>

*/
