/*

golang第三十四天:go语言中的原子计数器Atomic Counters
http://www.kuaizh.com/?p=723


go语言利用 sync/atomic 包实现一个被多个协程访问的原子计数器。
用于进行增或减的原子操作的函数名称都以"Add"为前缀，并后跟针对的具体类型的名称。

由于atomic.AddUint32函数和atomic.AddUint64函数的第二个参数的类型分别是uint32和uint64，
所以我们无法通过传递一个负的数值来减小被操作值。atomic.AddUint32(&ui32, ^uint32(-NN-1)) 其中NN代表了一个负整数

函数atomic.LoadInt32接受一个*int32类型的指针值，并会返回该指针值指向的那个值
有了"原子的"这个形容词就意味着，在这里读取value的值的同时，当前计算机中的任何CPU都不会进行其它的针对此值的读或写操作。

这样的约束是受到底层硬件的支持的。


*/

package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {

	// 利用一个无符号整数作为计数器(永远是正值)
	var ops uint64 = 0

	// 为了模拟并发更新，启动50个协程，每个协程每毫秒为计数器自增
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 使用AddUint64函数为计数器进行自增操作，向其传递计数器的内存地址作为第一个参数
				atomic.AddUint64(&ops, 1)
				// 允许其他协程进行处理
				runtime.Gosched()
			}
		}()
	}

	// 等待1秒以允许一些操作完成
	time.Sleep(time.Second)

	// 当其他协程正在更新的时候，为了安全使用计数器，我们通过 LoadUint64 释出一份当前值的拷贝到 opsFinal 中
	// 和上面一样，我们需要给这个函数传递计数器的内存地址
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

/*

运行结果显示我们执行了4000次操作：
func  go run atomic-counters.go
ops: 40200

*/
