/*

golang第五十二天:go语言中的命令行参数Command-Line Arguments
http://www.kuaizh.com/?p=769


命令行参数(命令行参数是指定程序运行参数的一个常见方式。例如，go run hello.go，程序 go 使用了 run 和 hello.go 两个参数。)


*/

package main

import "os"
import "fmt"

func main() {
	//os.Args 提供原始命令行参数访问功能。注意，切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有程序的的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	//你可以使用标准的索引位置方式取得单个参数的值。
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

/*

//要实验命令行参数，最好先使用 go build 编译一个可执行二进制文件

func  go build command-line-arguments.go
func  ./command-line-arguments a b c d
[./command-line-arguments a b c d]
[a b c d]
c

*/
