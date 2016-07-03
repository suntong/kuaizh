/*

golang第五十四天:go语言中的获取和设置环境变量Environment Variables
http://www.kuaizh.com/?p=773


Golang 要获取环境变量需要使用os包。导入"os"包，通过os包中的Getenv方法来获取。Setenv设置环境变量

*/

package main

import "os"
import "strings"
import "fmt"

func main() {
	//设置环境变量，然后读取设置的环境变量。读取环境变量是，如果key不存在，返回空字符串。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	//os.Environ列出所有的环境变量。返回一个key=value的分片。使用strings.Split来分割key，value。
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}

/*

执行结果：
func  go run environment-variables.go
FOO: 1
BAR:

//系统变量列表，依赖你的机器。

TERM_PROGRAM
PATH
SHELL
...

//如果我们开始就设置了BAR，也会读取到相应的值。
func  BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...

*/
