/*

golang第五十七天:go语言中的处理系统信号量Signals
http://www.kuaizh.com/?p=779


有时候我们需要处理系统的信号量，在实际项目中我们可能有下面的需求：

1、修改了配置文件后，希望在不重启进程的情况下重新加载配置文件；
2、当用 Ctrl + C 强制关闭应用后，做一些必要的处理；

这时候就需要通过信号传递来进行处理了。golang中对信号的处理主要使用os/signal包中的两个方法：一个是notify方法用来监听收到的信号；一个是 stop方法用来取消监听。

*/

package main

import "fmt"
import "time"
import "os"
import "os/signal"
import "syscall"

func main() {

	//Go 通过向一个通道发送 os.Signal 值来进行信号通知。我们将创建一个通道来接收这些通知（同时还创建一个用于在程序可以结束时进行通知的通道）。
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	//signal.Notify 注册这个给定的通道用于接收特定信号。
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//这个 Go 协程执行一个阻塞的信号接收操作。当它得到一个值时，它将打印这个值，然后通知程序可以退出。
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	//程序将在这里进行等待，直到它得到了期望的信号（也就是上面的 Go 协程发送的 done 值）然后退出。
	loop := true
	for loop {
		fmt.Println("awaiting signal")
		select {
		case res := <-done:
			loop = false
			fmt.Println(res)
		case <-time.After(time.Second * 1):
			fmt.Println("loop again")
		}
	}

	fmt.Println("exiting")
}

/*

//当我们运行这个程序时，它将一直等待一个信号。使用 ctrl-C（终端显示为 ^C），我们可以发送一个 SIGINT 信号，这会使程序打印 interrupt 然后退出。
C:\Users\Administrator>go run signal.go
awaiting signal
loop again
awaiting signal
loop again
awaiting signal
loop again
awaiting signal
loop again
awaiting signal
loop again
awaiting signal
loop again
awaiting signal
loop again
awaiting signal

interrupt
true
exiting

*/
