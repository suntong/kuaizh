/*

golang第五十五天:go语言中的启动进程以及调用shell等系统命令Spawning Processes
http://www.kuaizh.com/?p=775


有时候我们的go程序需要调用一些外部进程，比如系统命令，shell脚本等，使用exec包可以创建一个外部进程。

*/

package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {
	//我们将从一个简单的命令开始，没有参数或者输入，仅打印一些信息到标准输出流。exec.Command 函数帮助我们创建一个表示这个外部进程的对象。
	dateCmd := exec.Command("date")

	//.Output 是另一个帮助我们处理运行一个命令的常见情况的函数，它等待命令运行完成，并收集命令的输出。如果没有出错，dateOut 将获取到日期信息的字节。
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> date")
	fmt.Println(string(dateOut))

	//下面我们将看看一个稍复杂的例子，我们将从外部进程的stdin 输入数据并从 stdout 收集结果。
	grepCmd := exec.Command("grep", "hello")

	//这里我们明确的获取输入/输出管道，运行这个进程，写入一些输入信息，读取输出的结果，最后等待程序运行结束。
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	//上面的例子中，我们忽略了错误检测，但是你可以使用if err != nil 的方式来进行错误检查，我们也只收集StdoutPipe 的结果，但是你可以使用相同的方法收集StderrPipe 的结果。
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	//注意，当我们需要提供一个明确的命令和参数数组来生成命令，和能够只需要提供一行命令行字符串相比，你想使用通过一个字符串生成一个完整的命令，那么你可以使用 bash命令的 -c 选项：
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

/*

//生成的程序返回和我们直接通过命令行运行这些程序的输出是相同的。
func  go run spawning-processes.go
> date
Wed Oct 10 09:53:11 PDT 2012


> grep hello
hello grep


> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go

*/
