/*

golang第三十六天:go语言中的排序Sorting
http://www.kuaizh.com/?p=729


Go语言使用sort包对任意类型元素的集合进行排序。Go语言里有个专门排序的包"sort"。他可以给slice结构排序，内置了对string、等类型的排序。


*/

package main

import "fmt"
import "sort"

	


func main() {

	//对字符串序列排序
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)
	
	//对整数分片排序
    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

	//检查一个分片是否是排序好的。
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}

/*

	

执行结果
Strings: [a b c]
Ints:    [2 4 7]
Sorted:  true

*/
