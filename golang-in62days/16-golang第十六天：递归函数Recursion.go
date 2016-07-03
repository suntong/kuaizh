/*

golang第十六天:go语言中的递归函数Recursion
http://www.kuaizh.com/?p=579


Go编程语言支持递归，即要调用的函数本身。但是在使用递归时，程序员需要谨慎确定函数的退出条件，否则会造成无限循环。

递归函数是解决许多数学问题想计算一个数阶乘非常有用的，产生斐波系列等

package main
import "fmt"
func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n-1)
}
func main() {
  fmt.Println(fact(7))
}
$ go run recursion.go
5040

*/

package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func doPalindrome(s string) bool {
	if utf8.RuneCountInString(s) <= 1 {
		return true
	}
	word := strings.Trim(s, "\t \r\n\v")
	first, sizeOfFirst := utf8.DecodeRuneInString(word)
	last, sizeOfLast := utf8.DecodeLastRuneInString(word)
	if first != last {
		return false
	}
	return doPalindrome(word[sizeOfFirst : len(word)-sizeOfLast])
}

func IsPalindrome(word string) bool {
	s := ""
	s = strings.Trim(word, "\t \r\n\v")
	if len(s) == 0 || len(s) == 1 {
		return false
	}
	return doPalindrome(s)
}

func main() {
	args := os.Args[1:]
	for _, v := range args {
		ok := IsPalindrome(v)
		if ok {
			fmt.Printf("%s\n", v)
		}
	}
}
