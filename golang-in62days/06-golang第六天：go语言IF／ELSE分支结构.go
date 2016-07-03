/*

golang第六天:go语言IF/ELSE分支结构
http://www.kuaizh.com/?p=555


if条件语句不需要括号括起来，但是if和条件语句之间需要空格。if可以没有else语句。if语句可以进行预处理操作。
:~/practice$ vi ifelse.go

*/

package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	var precede = 5
	fmt.Println("precede is:", precede)
	if precede = precede * 2; precede < 0 {
		fmt.Println(precede, "is negative")
	} else if precede < 10 {
		fmt.Println(precede, "has 1 digit")
	} else {
		fmt.Println(precede, "has multiple digits")
	}
}

/*

:~/practice$ go run ifelse.go
7 is odd
8 is divisible by 4
9 has 1 digit
precede is: 5
10 has multiple digits
:~/practice$

*/
