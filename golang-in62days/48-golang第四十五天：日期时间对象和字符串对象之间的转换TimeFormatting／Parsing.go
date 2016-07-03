/*

golang第四十五天:go语言中的日期时间对象和字符串对象之间的转换Time Formatting / Parsing
http://www.kuaizh.com/?p=751


Go语言中，获取时间戳用time.Now().Unix()，获取时间对象用time.Now()，格式化时间用t.Format，解析时间用time.Parse。
在Go语言里，使用数字来表示格式：

月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
年　 06,2006
周几 Mon,Monday

时区时差表示 -07,-0700,Z0700,Z07:00,-07:00,MST

时区字母缩写 MST

*/

package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	//根据RFC3339来格式化时间
	t := time.Now()
	p(t.Format(time.RFC3339))

	//根据RFC3339来把字符串转化成时间对象
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	//根据时间模式来格式化时间。
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	//也可以根据标准的字符串格式化来输出时间字符串。
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	//格式不对是，Parse将会返回错误值。
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}

/*

执行结果：
2014-04-15T18:00:15-07:00
2012-11-01 22:08:41 +0000 +0000
6:00PM
Tue Apr 15 18:00:15 2014
2014-04-15T18:00:15.161182-07:00
0000-01-01 20:41:00 +0000 UTC
2014-04-15T18:00:15-00:00
parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": ...

*/
