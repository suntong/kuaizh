/*

golang第四天:go语言的常量
http://www.kuaizh.com/?p=551


常量是程序中最基础的元素，在定义之后就不能再重新赋值了。Go语言中的常量类型有布尔常量、整数常量、浮点数常量、 字符常量、字符串常量和复数常量 。
常量使用const来定义。
:~/practice$ vi const.go

*/

package main

import "fmt"
import "math"

const s string = "constant"

func main() {
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
	//n=1; cannot assign to n
}

/*

:~/practice$ go run const.go
constant
6e+11
600000000000
-0.28470407323754404


*/
