/*

golang第一天:golang hello world和for循环例子
http://www.kuaizh.com/?p=547


//hello world例子

~/practice$ touch main.go
~/practice$ vi main.go

package main
import "fmt"
func main() {
  fmt.Printf("hello, world\n")
}

~/practice$ go run main.go
hello, world

*/

/*

//for循环例子

*/

package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("hello world:%d\n", i)
	}
}

/*

~/practice$ go run main.go
hello, world
hello world:0
hello world:1
hello world:2
hello world:3
hello world:4
hello world:5
hello world:6
hello world:7
hello world:8
hello world:9

*/
