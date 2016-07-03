/*

golang第二十四天:go语言中的并发操作带缓冲的管道Channel Buffering
http://www.kuaizh.com/?p=669


默认channel是无缓冲的，相当于channel的容量是1，发送方发送完数据后，会一直阻塞直到接收方将数据取出。

如果channel带有缓冲区，发送方会一直发送数据到缓冲区；如果缓冲区已满，则发送方阻塞，只能在接收方取走数据后发送方才能从阻塞状态恢复。


*/

package main

import "fmt"

func main() {

	//容量为2的channel
	messages := make(chan string, 2)

	//我们可以发送2个字符串，而不会被阻塞
	messages <- "buffered"
	messages <- "channel"

	//接受这俩个值
	fmt.Println(<-messages)
	//再发送一个
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/*

C:\Users\Administrator>go run channel.go
buffered
channel
channel

*/
