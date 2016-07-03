/*

golang第三十天:go语言中的通过for-range循环读取channel
http://www.kuaizh.com/?p=681


for和range不只是可以遍历基础数据类型，还可以遍历管道channel。
for i := range ch { // ch关闭时，for循环会自动结束
    println(i)
}

*/

package main

import "fmt"

func main() {
	//queue发送俩个字符串。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	//for range遍历从queue中接收到的每一个元素，因为我们关闭了channel，所以便利完俩个元素后，循环终止。
	//如果我们没有close channel，第三次循环receive将会阻塞。
	for elem := range queue {
		fmt.Println(elem)
	}

}

/*

$ go run range-over-channels.go
one
two

我们可以close一个非空的channel，close后的channel依然可读取。

*/
