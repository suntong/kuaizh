/*

golang第九天:go语言中的分片Slices
http://www.kuaizh.com/?p=561


分片是go语言里面一种重要的数据类型，它也是一组元素的集合。和数组相比，分片提供更强大的接口。
和数组不同，分片只定义元素的类型，而没有指定分片所能包含的元素的数量。

- 创建分片，需要使用内置的make函数， 也可以指定创建分片时的初始分片长度。
- 内置方法append可以往分片添加一个或多个元素，append返回新的分片对象，而不是修改原来的分片对象。
- 内置方法copy可以从一个分片拷贝元素到另外一个分片。
- 分片支持分片截取[from:to]操作，包含from，不包含to。
- 我们可以通过:=定义和初始化分片。分片也可以嵌套。
- fmt.Println打印分片和打印数组输出的格式是一样的。

:~/practice$ vi slice.go

*/

package main

import "fmt"

func main() {
	s := make([]string, 3) //创建3个空字符串的分片。
	fmt.Println("emp:", s)
	//分片元素的读取和数组一样
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	//内置函数len也返回分片元素的个数。
	fmt.Println("len:", len(s))
	//内置方法append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	//内置方法copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	//分片截取[from:to]操作
	l := s[2:5]
	fmt.Println("sl1:", l)
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)
	//通过:=定义和初始化分片
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	//分片也可以嵌套
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

/*

:~/practice$ go run slice.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

*/
