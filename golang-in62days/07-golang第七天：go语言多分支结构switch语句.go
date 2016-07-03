/*

golang第七天:go语言多分支结构switch语句
http://www.kuaizh.com/?p=557


switch语句支持从一批可选的条件中，选择合适的语句执行。case里面可以包含多个可选的条件，用逗号分隔。
switch可以没有条件语句，这样的switch就想到if/else分支语句了。而且case可以调用外面的变量，不一定要是常量了。
:~/practice$ vi switch.go

*/

package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it’s the weekend")
	default:
		fmt.Println("it’s a weekday")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it’s before noon")
	default:
		fmt.Println("it’s after noon")
	}
}

/*

:~/practice$ go run switch.go
write 2 as two
it’s a weekday
it’s after noon
:~/practice$

*/
