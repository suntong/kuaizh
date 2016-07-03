/*

golang第二十六天:go语言中的阻塞式Select
http://www.kuaizh.com/?p=673


select是Go语言特有的操作，使用select我们可以同时在多个channel上进行发送/接收操作。

select {
case x := <- somechan:
    // ... 使用x进行一些操作
case y, ok := <- someOtherchan:
    // ... 使用y进行一些操作，
    // 检查ok值判断someOtherchan是否已经关闭
case outputChan <- z:
    // ... z值被成功发送到Channel上时
default:
    // ... 上面case均无法通信时，执行此分支
}

*/

package main

import "time"
import "fmt"

func main() {
	//我们在俩个channel上进行select
	c1 := make(chan string)
	c2 := make(chan string)

	//通过协程想channel发现数据
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	//我们同时在2个channel上面等待数据
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}

	}

}

/*

//我们先接受到one，然后接受到two
func  time go run select.go
received one
received two

总运行时间2秒钟，因为每个协程sleep了一秒。
real	0m2.245s

注意：
默认的select是阻塞式的。select会一直等待等到某个 case 语句完成， 也就是等到成功从 ch1 或者 ch2 中读到数据。 则 select 语句结束。

*/
