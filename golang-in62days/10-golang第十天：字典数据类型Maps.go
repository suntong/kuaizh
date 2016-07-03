/*

golang第十天:go语言中的字典数据类型Maps
http://www.kuaizh.com/?p=563


字典是Go语言内置的关联数据类型。因为数组是索引对应数组元素，而字典是键对应值。

创建一个字典可以使用内置函数make，"make(map[键类型]值类型)".
使用经典的"name[key]=value"来为键设置值.

用Println输出字典，会输出所有的键值对.
内置函数len返回字典的元素个数.
内置函数delete从字典删除一个键对应的值。
可以用 ":=" 同时定义和初始化一个字典。

根据键来获取值有一个可选的返回值，这个返回值表示字典中是否存在该键，如果存在为true，返回对应值，否则为false，
返回零值有的时候需要根据这个返回值来区分返回结果到底是存在的值还是零值，比如字典不存在键x对应的整型值，返回零值就是0，
但是恰好字典中有键y对应的值为0，这个时候需要那个可选返回值来判断是否零值。
:~/practice$ vi maps.go

*/

package main

import "fmt"

func main() {
	//make
	m := make(map[string]int)
	//根据key来设置值
	m["k1"] = 7
	m["k2"] = 13
	//打印字典
	fmt.Println("map:", m)
	//根据key来获取值
	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	//字典元素个数
	fmt.Println("len:", len(m))
	//delete元素
	delete(m, "k2")
	fmt.Println("map:", m)
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	//:=使用
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

/*

:~/practice$ go run maps.go
map: map[k1:7 k2:13]
v1:  7
len: 2
map: map[k1:7]
prs: false
map: map[foo:1 bar:2]
:~/practice$

*/
