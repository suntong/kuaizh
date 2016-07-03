/*

golang第三天:go语言的基本数据类型和变量
http://www.kuaizh.com/?p=549


go语言的基本数据类型包括字符串，整数，浮点数，布尔类型等。字符串可以通过"+"拼接。
:~/practice$ vi value.go


package main
import "fmt"
func main() {
  fmt.Println("go" + "lang")
  fmt.Println("1+1 =", 1+1)
  fmt.Println("7.0/3.0 =", 7.0/3.0)
  fmt.Println(true && false)
  fmt.Println(true || false)
  fmt.Println(!true)
}


:~/practice$ go run value.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
*/

/*

变量用来指向一个值或一个函数等。
:~/practice$ vi variables.go

*/

package main

import "fmt"

func main() {
	//通过var定义一个或多个变量
	var a string = "initial"
	fmt.Println(a)
	//可以一次定义多个变量.
	var b, c int = 1, 2
	fmt.Println(b, c)
	//没有定义类型的变量，可以通过初始值来确定变量的类型
	var d = true
	fmt.Println(d)
	//未初始化的变量，将采用默认的初始值，必然int类型，默认的初始值未0
	var e int
	fmt.Println(e)
	//":="用来定义和初始化一个变量，可以省略var和变量类型
	f := "short"
	fmt.Println(f)
}

/*

:~/practice$ go run variables.go
initial
1 2
true
0
short


*/
