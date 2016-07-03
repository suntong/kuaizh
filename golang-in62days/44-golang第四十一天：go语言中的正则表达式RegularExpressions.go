/*

golang第四十一天:go语言中的正则表达式Regular Expressions
http://www.kuaizh.com/?p=743


Golang中的正则表达式

对于 [a-z] 这样的正则表达式，如果要在 [] 中匹配 - ，可以将 - 放在 [] 的开头或结尾，例如 [-a-z] 或 [a-z-] 可以在 [] 中使用转义字符：\f、\t、\n、\r、\v、\377、\xFF、\x{10FFFF}、\、^、\$、.、*、+、\?、{、}、(、)、[、]、\|（具体含义见上面的说明）

如果在正则表达式中使用了分组，则在执行正则替换的时候，"替换内容"中可以使用 $1、${1}、$name、${name} 这样的"分组引用符"获取相应的分组内容。其中 $0 代表整个匹配项，$1 代表第 1 个分组，$2 代表第 2 个分组，......。

如果"分组引用符"是 $name 的形式，则在解析的时候，name 是取尽可能长的字符串，比如：$1x 相当于 ${1x}，而不是${1}x，再比如：$10 相当于 ${10}，而不是 ${1}0。

由于 $ 字符会被转义，所以要在"替换内容"中使用 $ 字符，可以用 \$ 代替。

上面介绍的正则表达式语法是"Perl 语法"，除了"Perl 语法"外，Go 语言中还有另一种"POSIX 语法"，"POSIX 语法"除了不能使用"Perl 类"之外，其它都一样。


*/

package main

import "bytes"
import "fmt"
import "regexp"

func main() {

	//这个测试一个字符串是否符合一个表达式。
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	//上面我们是直接使用字符串，但是对于一些其他的正则任务，你需要使用 Compile 一个优化的 Regexp 结构体。
	r, _ := regexp.Compile("p([a-z]+)ch")

	//这个结构体有很多方法。这里是类似我们前面看到的一个匹配测试。
	fmt.Println(r.MatchString("peach"))

	//这是查找匹配字符串的。
	fmt.Println(r.FindString("peach punch"))

	//这个也是查找第一次匹配的字符串的，但是返回的匹配开始和结束位置索引，而不是匹配的内容。
	fmt.Println(r.FindStringIndex("peach punch"))

	//Submatch 返回完全匹配和局部匹配的字符串。例如，这里会返回 p([a-z]+)ch 和 `([a-z]+) 的信息。
	fmt.Println(r.FindStringSubmatch("peach punch"))

	//类似的，这个会返回完全匹配和局部匹配的索引位置。
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	//带 All 的这个函数返回所有的匹配项，而不仅仅是首次匹配项。例如查找匹配表达式的所有项。
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	//All 同样可以对应到上面的所有函数。
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	//这个函数提供一个正整数来限制匹配次数。
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	//上面的例子中，我们使用了字符串作为参数，并使用了如 MatchString 这样的方法。我们也可以提供 []byte参数并将 String 从函数命中去掉。
	fmt.Println(r.Match([]byte("peach")))

	//创建正则表示式常量时，可以使用 Compile 的变体MustCompile 。因为 Compile 返回两个值，不能用常量。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	//regexp 包也可以用来替换部分字符串为其他值。
	fmt.Println(r.ReplaceAllString("a peach", ""))

	//Func 变量允许传递匹配内容到一个给定的函数中，
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

/*

执行结果：
true
true
peach
[0 5]
[peach ea]
[0 5 1 3]
[peach punch pinch]
[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
[peach punch]
true
p([a-z]+)ch
a
a PEACH

*/
