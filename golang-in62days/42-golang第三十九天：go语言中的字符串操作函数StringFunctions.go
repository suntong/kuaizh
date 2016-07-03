/*

golang第三十九天:go语言中的字符串操作函数String Functions
http://www.kuaizh.com/?p=739


标准库strings提供了很多操作字符串的函数。

*/

package main

import s "strings"
import "fmt"

//给fmt.Println函数提供一个别名，因为它用的比较频繁。
var p = fmt.Println

func main() {

	//strings包里面提供的函数，注意他们不是字符串对象的方法，而是strings包里面提供的内置函数。
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	//不在strings包里面，但是很有用的len方法以及根据索引获取字符的方法。
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}

/*

运行结果：
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
toLower:    test
ToUpper:    TEST



Len:  5
Char: 101

*/
