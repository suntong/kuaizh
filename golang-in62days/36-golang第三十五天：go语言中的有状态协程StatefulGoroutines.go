/*

golang第三十五天:go语言中的有状态协程Stateful Goroutines
http://www.kuaizh.com/?p=727


Go语言内存模型规定了在一个goroutine中一个变量的读取的情况下，确保能够观察到在其他另外goroutine中写入同样变量的值。

也就是说，如果在多个goroutine操作修改同一个变量状态情况下，Go内存模型能够保证一个goroutine对变量写入的数据能够被其他goroutine正常读取，

类似多线程编程中两个线程对同一个变量读写保证一样。

*/

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

//在这个例子中，我们的共享数据将会被一个协程管理，其他协程如果需要读写数据，则发送消息给这个管理协程。
//结构体readOp和writeOp表示一个读取消息和写入消息。
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	//记录操作次数
	var ops int64 = 0

	//其他的协程通过读写管道，发送读写消息。
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	//共享变量管理协程，共享变量作为协程的一个内部状态。协程循环读取读写管道，获取请求，处理数据，然后把相应处理结果写到结构体中的resp channel管道里面。
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	//启动100个协程读取数据，每个读操作先生成一个读请求结构体，发送结构体实例到reads channel管道。然后读取结构体中的响应管道获取数据。
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	//启动10个协程写入数据。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	//所有的协程执行1秒钟
	time.Sleep(time.Second)

	//读取操作次数
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
}

/*

//执行结果
ops: 807434

基于协程的操作比互斥锁性能低很多，在有些情况下，互斥锁不方便使用的情况下，可以使用基于协程的方式来实现。

*/
