/*

golang第二十三天:go语言中的并发操作管道Channels
http://www.kuaizh.com/?p=667


Goroutine和channel是Go在"并发"方面两个核心feature。

Channel是goroutine之间进行通信的一种方式，它与Unix中的管道类似。
channel是类型相关的，也就是一个channel只能传递一种类型。例如，上面的ch只能传递int。

在go语言中，有4种引用类型：slice，map，channel，interface。
通过make(chan val-type)来创建一个channel。通过channel <- 把数据发送到channel。通过<-channel来从channel中获取数据。


*/

package main

import "fmt"

func main() {
	//创建channel
	messages := make(chan string)

	//通过一个goroutine发送‘ping’到channel。
	go func() { messages <- "ping" }()

	//获取数据
	msg := <-messages
	fmt.Println(msg)
}

/*


func  go run channels.go
ping

当channel满或者channel为空时，默认的发送数据和接受数据会被阻塞。接收方会一直阻塞直到有数据到来。

*/
