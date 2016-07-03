/*

golang第三十九天:go语言中的集合操作函数Collection Functions
http://www.kuaizh.com/?p=737


我们常常需要操作集合，比如从集合中选择符合条件的所有元素，或者按照一定的规则转换集合中的所有元素，从而得到新的集合。
JAVA集合框架使用泛型提供通用的方法。go语言不支持泛型，在go语言里面，如果内置的函数不够，需要自定义函数来操作集合。
下面demo我们提供了字符串分片的一些集合操作函数，Index定位，Include包含,Filter过滤，map转换等等。


*/

package main

import "strings"
import "fmt"

//查找元素，返回第一个元素的位置，或者-1，如果不存在。

func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
}

/*

}

/*

    return -1
}

/*


//判断是否包含某个元素

func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

/*


//查询是否有元素满足给定的条件，条件为一个函数。

func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
}

/*

}

/*

    return false
}

/*


//检查是否所有的元素都满足指定的条件f

func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
}

/*

}

/*

    return true
}

/*


//查找所有满足条件的元素的集合。

func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
}

/*

}

/*

    return vsf
}

/*


//对集合的所有元素进行转换，生成另外一个集合

func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
}

/*

    return vsm
}

/*


	


func main() {
	//调用我们定义的集合操作函数
    var strs = []string{"peach", "apple", "pear", "plum"}
    fmt.Println(Index(strs, "pear"))
    fmt.Println(Include(strs, "grape"))
    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
}))

/*

	
    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
}))

/*

	
    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
}))

/*

	
	//上面都是使用内部匿名函数，也可以使用函数变量
    fmt.Println(Map(strs, strings.ToUpper))

	

}

/*


	

输出结果：
2
false
true
false
[peach apple pear]
[PEACH APPLE PEAR PLUM]

*/
