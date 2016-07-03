/*

golang第四十天:go语言中的字符串格式化String Formatting
http://www.kuaizh.com/?p=741


go语言以printf的形式提供了非常出色字符串格式化函数。下面是一些例子：
Go 的排版打印风格类似 C 的 printf 族但更丰富更通用。这些函数活在 fmt 包里，叫大写的名字：fmt.Printf，fmt.Fprintf，fmt.Sprintf 等等。
字串函数（Sprintf 等）返回字串，而不是填充给定的缓冲。

*/

package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {
	//可以用全拿格式 %v（代表 value）；结果和Print 与 Println 打印的完全一样。再有，此格式可打印任意值，包括数组，结构和映射。
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	//打印结构时，改进的格式 %+v 用结构的域名注释，对任意值格式 %#v 打印出完整的 Go 句法。
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)

	// %T，打印某值的类型
	fmt.Printf("%T\n", p)

	//%t打印布尔值
	fmt.Printf("%t\n", true)

	//%d打印10进制整数
	fmt.Printf("%d\n", 123)

	//%b打印2进制整数
	fmt.Printf("%b\n", 14)

	//%c打印整数对应的字符
	fmt.Printf("%c\n", 33)

	//%x打印16进制整数
	fmt.Printf("%x\n", 456)

	////%f打印标准的浮点数
	fmt.Printf("%f\n", 78.9)

	//%e and %E可行计数法
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	//%s打印标准字符串
	fmt.Printf("%s\n", "\"string\"")

	//%q给打印结果添加双引号
	fmt.Printf("%q\n", "\"string\"")

	//打印base-16的字符串
	fmt.Printf("%x\n", "hex this")

	//%p打印指针
	fmt.Printf("%p\n", &p)

	//打印整数，指定位数
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	//打印浮点数，指定位数和精度
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	//左对齐打印
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	//打印字符，控制长度和对齐方式，默认右对齐
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	//左对齐打印
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	//Sprintf返回字符串
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	//Fprintf可以指定输出终端。
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

/*

执行结果：
{1 2}
{x:1 y:2}
main.point{x:1, y:2}
main.point
true
123
1110
!
1c8
78.900000
1.234000e+08
1.234000E+08
"string"
"\"string\""
6865782074686973
0x42135100
|    12|   345|
|  1.20|  3.45|
|1.20  |3.45  |
|   foo|     b|
|foo   |b     |
a string
an error

*/
