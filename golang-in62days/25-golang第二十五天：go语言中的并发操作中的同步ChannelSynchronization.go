/*

golang第二十五天:go语言中的并发操作中的同步Channel Synchronization
http://www.kuaizh.com/?p=671


main goroutine通过"<-c"来等待sub goroutine中的"完成事件"，sub goroutine通过close channel促发这一事件。

当然也可以通过向Channel写入一个bool值的方式来作为事件通知。main goroutine在channel c上没有任何数据可读的情况下会阻塞等待。


*/

package main

import "fmt"
import "time"

// This is the function we’ll run in a goroutine. The done channel will be used to notify another goroutine that this function’s work is done.
//worker函数在协程里面运行，参数done channel用来通知主协程，该函数已经执行完成。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//发送数据true到channel
	done <- true
}

func main() {
	//启动协程，传递channel参数
	done := make(chan bool, 1)
	go worker(done)

	//main协程阻塞，等待worker协程完成。
	<-done
}

/*

输出结果：
working...done

如果删除了<-done，worker协程还没有启动，程序就会退出。

*/
