/*

golang第二十七天:go语言中的超时实现方式Timeouts
http://www.kuaizh.com/?p=675


Go 语言的 channel 本身是不支持 timeout 的，所以一般实现 channel 的读写超时都采用 select

select {
case <-c:
case <-time.After(time.Second):
}

*/

package main

import "time"
import "fmt"

func main() {

	//假设我们现在调用一个外部资源，俩分钟后返回结果到c1
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	//select等待结果，<-Time.After当c1等待结果超过1秒是，触发超时。
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	//如果把超时设置为3秒，则可以成功等待c2的结果。
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

}

/*

输出结果，第一个超时，没有结果，第二个成功返回，不触发超时语句。

$ go run timeouts.go
timeout 1
result 2

*/
