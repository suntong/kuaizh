/*

golang第十二天:go语言中的函数Functions定义和调用
http://www.kuaizh.com/?p=571


函数是Go里面的核心设计，它通过关键字func来声明。

func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
//这里是处理逻辑代码
//返回多个值
return value1, value2
}

1. 关键字func用来声明一个函数funcName
2. 函数可以有一个或者多个参数，每个参数后面带有类型，通过,分隔
3. 函数可以返回多个值
4. 上面返回值声明了两个变量output1和output2，如果你不想声明也可以，直接就两个类型
5. 如果只有一个返回值且不声明返回值变量，那么你可以省略 包括返回值 的括号
6. 如果没有返回值，那么就直接省略最后的返回信息
7. 如果有返回值， 那么必须在函数的外层添加return语句

*/

package main

import "fmt"

//求和函数
func plus(a int, b int) int {
	return a + b
}

//如果多个参数类型一样，可以省略前面参数的类型。
func plusPlus(a, b, c int) int {
	return a + b + c
}
func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

/*

$ go run functions.go
1+2 = 3
1+2+3 = 6

*/
