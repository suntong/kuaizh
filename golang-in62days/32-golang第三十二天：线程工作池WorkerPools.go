/*

golang第三十二天:go语言中的线程工作池Worker Pools
http://www.kuaizh.com/?p=719


使用协程和goroutine管道channel，可以方便的实现线程工作池。
比如线程池开启number个线程，每个线程从任务队列中取出一个任务执行，执行完成后取下一个任务。
可以把任务放到channel里，每个线程不停的从channel中取出任务执行，并把执行结果写入另一个channel，当任务channel关闭时，线程就会退出for range循环。

*/

package main

import "fmt"
import "time"

//工作函数，我们将通过协程启动多个实例，他们读取jobs channel里面的数据，进行处理，然后把结果发送到result channel里面。
//我们使用sleep来模拟比较耗时的任务处理。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}

}

func main() {
	//定义俩个管道，任务channel和结果channel
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//启动3个协程，因为jobs里面没有数据，他们都会阻塞。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//发送9个任务到jobs channel，然后关闭管道channel。
	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	close(jobs)

	//然后收集所有的结果。
	for a := 1; a <= 9; a++ {
		<-results
	}

}

/*


//我们发现jobs里面的任务，被三个协程同步处理了。

func  time go run worker-pools.go
worker 1 processing job 1
worker 2 processing job 2
worker 3 processing job 3
worker 1 processing job 4
worker 2 processing job 5
worker 3 processing job 6
worker 1 processing job 7
worker 2 processing job 8
worker 3 processing job 9

程序运行后显示9个任务都被这些worker执行了。一共9秒的任务，程序一共运行了大概3秒就完成了，这是因为3个协程在并发操作。


*/
