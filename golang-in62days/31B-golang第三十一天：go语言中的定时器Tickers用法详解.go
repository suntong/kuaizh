/*

golang第三十一天:go语言中的定时器Tickers用法详解
http://www.kuaizh.com/?p=717


Timers解决在多久的将来执行一次的问题，tickers解决按时间间隔重复执行的问题。
tickers的用法和timers很相似。


*/

package main

import "time"
import "fmt"

func main() {

	//time.NewTicker生成一个ticket，它包含一个管道channel C，然后每个相应的时间间隔，会向管道发送数据。
	//我们使用for range遍历管道，就实现了间隔时间定时执行的问题。
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}

	}()

	//和timers一样，tickes也可以被停止，停止后，管道就不会接受值了
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

/*


tickets被触发3次后停止。

func  go run tickers.go
Tick at 2012-09-23 11:29:56.487625 -0700 PDT
Tick at 2012-09-23 11:29:56.988063 -0700 PDT
Tick at 2012-09-23 11:29:57.488076 -0700 PDT
Ticker stopped

*/
