/*

golang第三十三天:go语言中的资源利用限制以及QOS:Rate Limiting
http://www.kuaizh.com/?p=721


Rate limiting 在 Web 架构中非常重要，是互联网架构可靠性保证重要的一个方面。
从最终用户访问安全的角度看，设想有人想暴力碰撞网站的用户密码，或者有人攻击某个很耗费资源的接口，
或者有人想从某个接口大量抓取数据。应该增加 Rate limiting，做请求频率限制。
Rate limiting可以很好的控制资源的使用情况，以及管理QOS。go通过goroutines, channels, and tickers可以很好的实现资源监控。

*/

package main

import "time"
import "fmt"

func main() {

	//限制处理请求数量。把请求放到一个管道channel里面进行处理。
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}

	close(requests)

	//limiter channel每200毫秒接受一个请求，200毫秒是我们处理请求所需要的时间。
	limiter := time.Tick(time.Millisecond * 200)

	//当我们处理请求时停止接受请求。这样通过定时器我们每200毫秒请求一次。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//有时候我们可以接受短时间内的请求超过我们的限制，我们可以通过缓冲管道来处理。
	//带3个缓冲的channel。
	burstyLimiter := make(chan time.Time, 3)

	//填满缓冲请求数量，表示可以立刻处理，不需要等待200毫秒才往缓冲管道里面填数据。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//每200毫秒添加一个请求到缓冲管道channel，直到填满后才阻塞，即停止发送数据到缓冲channel。
	//模拟每200毫秒允许处理一个请求。
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	//现在同时进来5个请求，因为缓冲队列里面已经填满了3个缓冲数据，所以这5个数据中的前面三个会马上被处理。
	//然后就会没200毫秒处理一个请求。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}

/*

$ go run rate-limiting.go
//每200毫秒处理一次
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC

//前面3个迅速处理，后面2个每200毫秒处理一次
request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC


*/
