/*

golang第五天:go语言唯一的循环结构FOR循环
http://www.kuaizh.com/?p=553


go语言只有一种循环结构，那就是for循环。
:~/practice$ vi for.go

*/

package main

import "fmt"

func main() {
	//最基本的循环结构，实现了while循环的功能。初始化条件在循环外面，条件在for语句里面，条件递增在循环代码里面。
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//典型的for循环结构。
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	//没有带条件的死循环结构，在循环代码里面判断是否需要跳出循环。
	for {
		fmt.Println("loop")
		break
	}
}

/*

:~/practice$ go run for.go
1
2
3
7
8
9
loop

*/
