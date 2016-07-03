/*

golang第四十四天:go语言中的unix时间戳Epoch
http://www.kuaizh.com/?p=749


通常我们需要获取unix时间戳，即从1970年1月1日（UTC/GMT的午夜）开始所经过的秒数，毫秒，纳秒等。
一个小时表示为UNIX时间戳格式为：3600秒；一天表示为UNIX时间戳为86400秒

*/

package main

import "fmt"
import "time"

func main() {

	//通过time的Unix()或者UnixNano()来获取
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	//没有毫秒方法，需要根据纳秒来转换。
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	//也可以把整形的毫秒或者纳秒转换成对应的时间对象。
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

/*

$ go run epoch.go
2012-10-31 16:13:58.292387 +0000 UTC
1351700038
1351700038292
1351700038292387000
2012-10-31 16:13:58 +0000 UTC
2012-10-31 16:13:58.292387 +0000 UTC

*/
