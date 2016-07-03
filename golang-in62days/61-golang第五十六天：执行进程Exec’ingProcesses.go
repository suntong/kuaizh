/*

golang第五十六天:go语言中的执行进程Exec’ing Processes
http://www.kuaizh.com/?p=777


有时候，我们只想用其他的（也许是非 Go 程序）来完全替代当前的 Go 进程。这时候，我们可以使用经典的 exec方法的 Go 实现。


*/

package main

import "syscall"
import "os"
import "os/exec"

func main() {

	//在我们的例子中，我们将执行ls命令。Go需要提供我们需要执行的可执行文件的绝对路径，所以我们将使用exec.LookPath来得到它（大概是/bin/ls）。
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	//Exec需要的参数是切片的形式的（不是放在一起的一个大字符串）。我们给ls一些基本的参数。注意，第一个参数需要是程序名。
	args := []string{"ls", "-a", "-l", "-h"}

	//Exec同样需要使用环境变量。这里我们仅提供当前的环境变量。
	env := os.Environ()

	//这里是os.Exec调用。如果这个调用成功，那么我们的进程将在这里被替换成 /bin/ls -a -l -h 进程。如果存在错误，那么我们将会得到一个返回值。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

/*

//当我们运行程序师，它会替换为 ls。
func  go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

注意 Go 并不提供一个经典的 Unix fork 函数。通常这不是个问题，因为运行 Go 协程，生成进程和执行进程覆盖了fork 的大多数使用用场景。

*/
