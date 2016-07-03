/*

golang第十三天:go语言中的多返回值函数Functions定义和调用
http://www.kuaizh.com/?p=573


Go语言内置支持多返回值，比如一个函数同时返回结果和错误信息。

*/

package main

import "fmt"

// 这个函数的返回值为两个int
func vals() (int, int) {
	return 3, 7
}

//交换两个数值,如果在其它语言中，我们第一想到的肯定是建立一个中间变量做交换。但在Go使用返回多个值的功能很容易的就实现了。
func swap(a int, b int) (int, int) {
	return b, a
}
func main() {
	// 获取函数的两个返回值
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	//声明的变量没有使用的话，编译无法通过
	// 如果你只对多个返回值里面的几个感兴趣
	// 可以使用下划线(_)来忽略其他的返回值
	_, c := vals()
	fmt.Println(c)
}

/*

输出结果为
3
7
7


*/
