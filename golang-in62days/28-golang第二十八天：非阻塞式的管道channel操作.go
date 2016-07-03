/*

golang第二十八天:go语言中的非阻塞式的管道channel操作
http://www.kuaizh.com/?p=677


对于Channel基本的收发是阻塞的。但是，我们可以使用 select 加 default 来实现非阻塞的收发，或者是多路select操作。


*/

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	//这里实现了非阻塞的接收。如果 messages 中有值，那么select将会先处理这个case，如果messages中没有值，那么select将会直接处理default分支的流程。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	//非阻塞发送数据到channel
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	//在多个管道channel上面实现非阻塞式读取数据，当messages和signals上面都没有数据时，马上执行default从句。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}

/*

$ go run non-blocking-channel-operations.go
no message received
no message sent
no activity

如果case都阻塞，则走default，如果无default，则阻塞在case
default中可以不读写任何通道，那么只要default提供不阻塞的出路，就相当于实现了对case的无阻塞尝试读写

*/
