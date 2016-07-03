/*

golang第三十一天:go语言中的定时器Timers用法详解
http://www.kuaizh.com/?p=715


我们常常需要在将来的某个时刻执行一些操作，或者摸个操作需要间隔多久重复执行。

go语言通过timer和ticket来实现这俩个功能。

*/

package main

import "time"
import "fmt"

func main() {

	//timer表示在未来一个单独的事件。你告诉timer需要等待多久，然后它生成一个channel，这个channel在到达时间后会触发。
	timer1 := time.NewTimer(time.Second * 2)

	//<-timer1.C 阻塞，直到定时器管道C被触发。相当于定时器网管道channel C写入了某个值，或者直接关闭管道。
	<-timer1.C
	fmt.Println("Timer 1 expired")

	//如果只是想等待多久再执行，也可以用time.Sleep。用timer的好处是，再触发之前，可以取消定时操作。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

}

/*

//第一个定时器等待2秒后执行，第二个定时器在触发之前就被我妈取消了。

$ go run timers.go
Timer 1 expired
Timer 2 stopped

下面三段代码(A,b,C)的功能都是在5分钟后执行指定的函数的golang代码
// (A)

time.AfterFunc(5 * time.Minute, func() {
    fmt.Printf("expired")
})


// (B) create a Timer object
timer := time.NewTimer(5 * time.Minute)
<-timer.C
fmt.Printf("expired")

// (C) time.After() 直接返回 timer.C
<-time.After(5 * time.Minute)
fmt.Printf("expired")

*/
