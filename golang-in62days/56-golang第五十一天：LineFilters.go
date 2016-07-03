/*

golang第五十一天:go语言中的Line Filters
http://www.kuaizh.com/?p=767


行过滤是一种常用的编程方式，它读取标准输入，进行转换，然后把结果打印到标准输出。比如linux的grep和sed。

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//包装标准缓冲输入到scanner，scanner提供了一个便利的Scan()方法，可以方便的读取下一行标识，默认是下一行数据。
	scanner := bufio.NewScanner(os.Stdin)

	//返回下一个输入
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		//输出转换后的字符串。
		fmt.Println(ucl)
	}

	//检查遍历过程中是否有错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

/*

输出结果：
func  echo 'hello'   > /tmp/lines
func  echo 'filter' >> /tmp/lines


func  cat /tmp/lines | go run line-filters.go
HELLO
FILTER

*/
