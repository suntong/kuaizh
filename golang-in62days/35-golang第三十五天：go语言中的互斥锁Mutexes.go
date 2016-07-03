/*

golang第三十五天:go语言中的互斥锁Mutexes
http://www.kuaizh.com/?p=725


互斥锁是传统的并发程序对共享资源进行访问控制的主要手段。它由标准库代码包sync中的Mutex结构体类型代表。

sync.Mutex类型（确切地说，是*sync.Mutex类型）只有两个公开方法——Lock和Unlock。顾名思义，前者被用于锁定当前的互斥量，而后者则被用来对当前的互斥量进行解锁。

类型sync.Mutex的零值表示了未被锁定的互斥量。通过互斥锁可以确保多个协程安全的访问数据。

对于同一个互斥锁的锁定操作和解锁操作总是应该成对的出现。如果我们锁定了一个已被锁定的互斥锁，

那么进行重复锁定操作的Goroutine将会被阻塞，直到该互斥锁回到解锁状态

*/

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//需要同步的共享数据
	var state = make(map[int]int)

	//互斥锁确保协程同步访问数据
	var mutex = &sync.Mutex{}

	//ops用来计算我们访问数据的次数
	var ops int64 = 0

	//我们启动100个协程同步的读取状态数据。
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				//我们选中一个key，锁住map，然后读取map数据，再解锁。计数器加1.
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				//允许其他协程使用cpu
				runtime.Gosched()
			}
		}()
	}

	//启动10个协程写数据到map。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	//读写协程执行1秒钟。
	time.Sleep(time.Second)

	//读取计数器
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	//获取锁，打印map的数据。
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

/*

//执行结果3,500,000

ops: 3598302
state: map[1:38 4:98 2:23 3:85 0:44]

*/
