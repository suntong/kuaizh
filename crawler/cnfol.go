package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	g, e := goquery.NewDocument("http://gold.3g.cnfol.com/")
	if e != nil {
		fmt.Println(e)
	}
	g.Find("ul").Eq(6).Find("a").
		Each(func(i int, content *goquery.Selection) {
			text := content.Text()
			a, _ := content.Attr("href")
			fmt.Printf("%s\n%s\n\n", text, a)
		})
}

/*

黄金跌了也不怕：若经济危机再临 金价将涨数倍
http://3g.cnfol.com/gold/201605/22790868.shtml

为什么大佬和散户都喜欢黄金？
http://3g.cnfol.com/gold/201605/22786069.shtml

破冰点金：晚盘黄金如何操作？
http://3g.cnfol.com/gold/201605/22790816.shtml

赵相宾：G7恐燃汇市地震 英国脱欧公投临近 黄金再主沉浮
http://3g.cnfol.com/gold/201605/22790814.shtml

*/
