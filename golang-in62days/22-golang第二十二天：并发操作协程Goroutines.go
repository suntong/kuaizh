/*

golang第二十二天:go语言中的并发操作协程Goroutines
http://www.kuaizh.com/?p=665


在go语言中，使用goroutine来实现并发程序,它是一种轻量级的线程。
go 函数(s)，就是把函数放在一个Goroutine协程里面同步指向。

*/

package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}

}

func main() {
	//采用同步方式调用f函数。
	f("direct")

	//采用协程异步执行。
	go f("goroutine")

	//通过调用里面函数的方式来启动协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//上面俩个函数在协程里面异步执行，我们通过键盘输入等他他们执行完成。
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

/*

执行结果：
func  go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2

done

要理解这个事儿首先得了解操作系统是怎么玩线程的。一个线程就是一个栈加一堆资源。操作系统一会让cpu跑线程A，一会让cpu跑线程B，靠A和B的栈来保存A和B的执行状态。

每个线程都有他自己的栈。但是线程又老贵了，花不起那个钱，所以go发明了goroutine。大致就是说给每个goroutine弄一个分配在heap里面的栈来模拟线程栈。

比方说有3个goroutine，A,B,C，就在heap上弄三个栈出来。然后Go让一个单线程的scheduler开始跑他们仨。相当于 { A(); B(); C() }，连续的，串行的跑。

和操作系统不太一样的是，操作系统可以随时随地把你线程停掉，切换到另一个线程。这个单线程的scheduler没那个能力啊，他就是user space的一段朴素的代码，

他跑着A的时候控制权是在A的代码里面的。A自己不退出谁也没办法。所以A跑一小段后需要主动说，老大（scheduler），我不想跑了，帮我把我的所有的状态保存在我自己的栈上面，让我歇一会吧。

这时候你可以看做A返回了。A返回了B就可以跑了，然后B跑一小段说，跑够了，保存状态，返回，然后C再跑。C跑一段也返回了。

这个关键就在于每个goroutine跑一跑就要让一让。一般支持这种玩意（叫做coroutine）的语言都是让每个coroutine自己说，我跑够了，换人。
goroutine比较文艺的地方就在于，他可以来帮你判断啥时候"跑够了"。其中有一大半就是靠的你说的"异步并发"。

go把每一个能异步并发的操作，像文件访问啦，网络访问啦之类的都包包好，包成一个看似朴素的而且是同步的"方法"，比如string readFile。

但是神奇的地方在于，这个方法里其实会调用"异步并发"的操作，比如某操作系统提供的asyncReadFile。这种异步方法都是很快返回的。

*/
