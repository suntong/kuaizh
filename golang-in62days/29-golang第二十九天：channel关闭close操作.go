/*

golang第二十九天:go语言中的channel关闭close操作
http://www.kuaizh.com/?p=679


关闭管道channel表示不会再往管道里面发送数据了。管道关闭后，等待接受数据的协程将会跳出阻塞。


*/

package main

import "fmt"

//主协程通过jobs管道来发送数据给worker协程。当主协程发送完后，close关闭，告诉worker协程没有待处理数据了。

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	//worker协程，反复从jobs管道接受数据。j, more := <-jobs，当jobs关闭后，并且jobs channel没有数据时，more返回false。
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {

				fmt.Println("received all jobs")
				done <- true
				return
			}

		}

	}()

	//发送3个任务，然后close channel。
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the synchronization approach we saw earlier.
	<-done
}

/*

$  go run closing-channels.go
sent job 1
received job 1
sent job 2
received job 2
sent job 3
received job 3
sent all jobs
received all jobs

Channel用作信号(Signal)等待一个事件，有时候通过close一个Channel就足够了。例如：

*/

/*

--------------------------------------------------

package main

import "fmt"


func main() {
  fmt.Println("Begin doing something!")
  c := make(chan bool)
  go func() {
    fmt.Println("Doing something...")
    close(c)
  }()

  data, more <-c
  fmt.Println("Done!", more)
}


这里main goroutine通过"<-c"来等待sub goroutine中的"完成事件"，sub goroutine通过close channel促发这一事件。当然也可以通过向Channel写入一个bool值的方式来作为事件通知。

main goroutine在channel c上没有任何数据可读的情况下会阻塞等待。

*/
