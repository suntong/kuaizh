/*

golang第十四天:go语言中的可变长参数函数Variadic Functions定义和调用
http://www.kuaizh.com/?p=575


支持可变长参数列表的函数可以支持任意个传入参数，比如fmt.Println函数就是一个支持可变长参数列表的函数。

对于可变长参数，go会创建一个slice，用来存放传入的可变参数，那么，如果创建一个slice，例如a，然后以...这种方式传入，

go会不会还会新建一个slice，将a的数据全部拷贝一份过去？答案是不会。

*/

package main

import "fmt"

//可以传入任意数量的整型参数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func t(args ...int) {
	fmt.Printf("%p\n", args)
}

func main() {
	// 支持可变长参数的函数调用方法和普通函数一样
	// 也支持只有一个参数的情况
	sum(1, 2)
	sum(1, 2, 3)
	// 如果你需要传入的参数在一个切片中，像下面一样
	// "func(slice...)"把切片打散传入
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	//打印参数地址%p
	a := []int{1, 2, 3}
	t(a...)
	fmt.Printf("%p\n", a)
}

/*

输出结果：
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10
0x1052e120
0x1052e120

*/
