/*

golang第三十七天:go语言中的利用函数排序Sorting by Functions
http://www.kuaizh.com/?p=731


有时候我们需要使用我们自己的逻辑来对集合进行排序，而不是安装集合元素的自然规则排序。

比如有时候我们希望根据字符串字符数量从多到少排序，而不是根据字符的规则排序。

这是可以利用我们自己定义的函数来排序。

*/

package main

import "sort"
import "fmt"

//定义一个ByLength类型，他就是字符串分片的别名。
type ByLength []string

//我们ByLength类型需要实现sort.Interface接口，包括Len，Less 和Swap方法。
//Len() 求长度、 Less(i,j) 比较第 i 和 第 j 个元素大小的函数、 Swap(i,j) 交换第 i 和第 j 个元素的函数。
//Len和Swap和平时一样，不需要修改。Less将会加入我们自己的排序逻辑。
//这个例子中，我们将会根据字符数量来排序，所以我们使用len(s[i])和len(s[j])。

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

//使用sort.Sort来排序，sort会用到我们自定义的类型。

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

/*

输入结果：
[kiwi peach banana]

*/
