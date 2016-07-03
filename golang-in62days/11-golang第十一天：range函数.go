/*

golang第十一天:go语言中的range函数
http://www.kuaizh.com/?p=569


range是go语言系统定义的一个函数。

函数的含义是在一个数组、分片、映射等数据结构中遍历每一个值，返回该值的下标值和此处的实际值。

假如说a[0]=10，则遍历到a[0]的时候返回值为0，10两个值。

*/

package main

import "fmt"

func main() {
	//使用range去遍历一个切片中的数值，然后求和。对于数组也是这样的操作的。
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//range 会同时返回索引和值。上面一个例子总我们不需要值的索引所以我们使用下划线将索引值废弃。现在我们打印索引值。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//遍历映射射，返回对应的键和值。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//对字符串使用range将会按 Unicode迭代返回字符串中的字符。返回的第一个值是 rune 的索引值，第二个值是 rune 值本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
