/*

golang第八天:go语言中的数组array
http://www.kuaizh.com/?p=559


1，go语言中，数组和其他语言C,JAVA一样，都是指定长度的同意类型的元素构成。
2，通过下标调用数组中的特定的元素。
3，内置函数len(a)返回数组的长度。
4，二维数组和多维数组也和C，JAVA一样。
5，fmt.Println打印数组是，格式是[元素1 元素2 。。。。。]
:~/practice$ vi array.go

*/

package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

/*

:~/practice$ go run array.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]
:~/practice$

*/
