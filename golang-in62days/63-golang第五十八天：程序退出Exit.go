/*

golang第五十八天:go语言中的程序退出Exit
http://www.kuaizh.com/?p=781


使用os.Exit来退出程序，返回状态码。

*/

package main

import "fmt"
import "os"

func main() {

	//当程序退出时，defer语句不会被执行。
	defer fmt.Println("!")
	//退出程序，返回状态码3
	os.Exit(3)
}

/*

执行结果：
func  go run exit.go
exit status 3

*/
