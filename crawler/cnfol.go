package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

/*

http://gold.3g.cnfol.com/

    <section> 
        <ul class="HotNew Mg10">
                        <li>
                <span class="HNTitTu"></span>
                <a href="http://3g.cnfol.com/gold/201605/22825169.shtml">美国耐用品经济数据不乐观 黄金止跌回稳</a>            
            </li>
                        <li>
                <span class="HNTitTu"></span>
                <a href="http://3g.cnfol.com/gold/201605/22825335.shtml">看空是错！下跌恰是黄金牛市的开始</a>            
            </li>
                        <li>
                <span class="HNTitTu"></span>
                <a href="http://3g.cnfol.com/gold/201605/22832590.shtml">ICN:黄金若下破1219.00位置将加速下跌</a>            
            </li>
                        <li>
                <span class="HNTitTu"></span>
                <a href="http://3g.cnfol.com/gold/201605/22832591.shtml">鸽派耶伦变调门加息预期大增 黄金加速探底</a>            
            </li>
                    
        </ul>
    </section>

*/

func main() {
	g, e := goquery.NewDocument("http://gold.3g.cnfol.com/")
	if e != nil {
		fmt.Println(e)
	}
	g.Find("ul.HotNew").Find("a").
		Each(func(i int, content *goquery.Selection) {
		text := content.Text()
		a, _ := content.Attr("href")
		fmt.Printf("%s\n%s\n\n", text, a)
	})
}

/*

美国耐用品经济数据不乐观 黄金止跌回稳
http://3g.cnfol.com/gold/201605/22825169.shtml

看空是错！下跌恰是黄金牛市的开始
http://3g.cnfol.com/gold/201605/22825335.shtml

ICN:黄金若下破1219.00位置将加速下跌
http://3g.cnfol.com/gold/201605/22832590.shtml

鸽派耶伦变调门加息预期大增 黄金加速探底
http://3g.cnfol.com/gold/201605/22832591.shtml

*/
