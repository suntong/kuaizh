/*

golang第四十七天:go语言中的字符串转换成数字Number Parsing
http://www.kuaizh.com/?p=755


把字符串转换成数字是最常用的功能。go语言的strconv包提供了把字符串转换成数字的功能。

*/

package main

import "strconv"
import "fmt"

func main() {

	//ParseFloat转换成浮点数。
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	//ParseInt转换成整数
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	//ParseInt会自动识别16进制字符串
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	//转换成无符号整数
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	//转换10进制整数
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	//返回错误
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

/*

执行结果：
1.234
123
456
789
135
strconv.ParseInt: parsing "wat": invalid syntax

*/
