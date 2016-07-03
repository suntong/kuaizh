/*

golang第四十三天:go语言中的日期时间操作Time
http://www.kuaizh.com/?p=747


time处理包括俩个方面：时间点（某一时刻）和时长（某一段时间）
编程离不开时间，时间管理，严格的说分成两块，一个是当前的时刻，对应的是一个点，还有是一段时间间隔。

*/

package main

import "fmt"
import "time"

func main() {
	p := fmt.Println
	//当前时间
	now := time.Now()
	p(now)

	//通过年月日时分秒生成时间对象
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	//获取年月日时分秒
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	//获取星期几
	p(then.Weekday())

	//时间比较
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	//时间相减，获取时长
	diff := now.Sub(then)
	p(diff)

	//时长换算成时分秒
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	//日期加减时长
	p(then.Add(diff))
	p(then.Add(-diff))
}

/*

$ go run time.go
2012-10-31 15:50:13.793654 +0000 UTC
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
25891h15m15.142266763s
25891.25420618521
1.5534752523711128e+06
9.320851514226677e+07
93208515142266763
2012-10-31 15:50:13.793654 +0000 UTC
2006-12-05 01:19:43.509120474 +0000 UTC

*/
