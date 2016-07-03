/*

golang第十七天:go语言中的指针Pointers
http://www.kuaizh.com/?p=592


&符号的意思是对变量取地址，如：变量a的地址是&a
*符号的意思是对指针取值，如:*&a，就是a变量所在地址的值，当然也就是a的值了

Golang 保留着C中值和指针的区别，但是对于指针繁琐用法进行了大量的简化，引入引用的概念。

所以在 Golang 中，你几乎不用担心会因为直接操作内寸而引起各式各样的错误。

GO中的map和slice是天然的引用类型，什么意思？
就是无论你怎么赋值，在GO的内部编译执行的时候，都是指针传递，并不会发生实质的内容拷贝

*/

package main

import "fmt"

//传递值，调用的时候，拷贝值到一个新的内存地址。申请了俩块内存。
func zeroval(ival int) {
	ival = 0
}

//传递指针，指向原来的内存地址。
func zeroptr(iptr *int) {
	//*iptr表示修改内存地址所对应的值对象。
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	//原来的i变量的值没有改变
	fmt.Println("zeroval:", i)
	//传递内存地址
	zeroptr(&i)
	//原来的i变量的值已经改变
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}

/*

$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100

*/
