/*

golang第三十八天:go语言中的异常Panic
http://www.kuaizh.com/?p=733


Go语言追求简洁优雅，所以，Go语言不支持传统的 try...catch...finally 这种异常，因为Go语言的设计者们认为，将异常与控制结构混在一起会很容易使得代码变得混乱。
因为开发者很容易滥用异常，甚至一个小小的错误都抛出一个异常。在Go语言中，使用多值返回来返回错误。
不要用异常代替错误，更不要用来控制流程。在极个别的情况下，也就是说，遇到真正的异常的情况下（比如除数为0了）。
才使用Go中引入的Exception处理：defer, panic, recover。
这几个异常的使用场景可以这么简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

panic是用来表示非常严重的不可恢复的错误的。在Go语言中这是一个内置函数，接收一个interface{}类型的值（也就是任何值了）作为参数。
panic的作用就像我们平常接触的异常。不过Go可没有try...catch，所以，panic一般会导致程序挂掉（除非recover）。panic就是这么简单。抛出个真正意义上的异常。



*/

package main

import "os"

	


func main() {

	//panic抛出一个不可预见的异常。
    panic("a problem")

	//通常的用法是当我们调用一个函数时，函数返回了错误，而我们不知道怎么处理这个错误。这时可以使用panic抛出异常。
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
}

/*

}

/*

	
//运行这个程序，将会抛出异常，以及协程堆栈信息。返回非0的状态码。
//程序在第一个panic的时候就退出了。不会再往下面执行。
func  go run panic.go
panic: a problem


goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

*/
