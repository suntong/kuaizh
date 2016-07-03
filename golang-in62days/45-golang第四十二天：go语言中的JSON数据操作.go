/*

golang第四十二天:go语言中的JSON数据操作
http://www.kuaizh.com/?p=745


Go语言转换JSON数据真是非常的简单。Go按照RFC 4627的标准实现了一个json编解码的标准库。

Marshal 用于将对象序列化到json对象中，Unmarshal用于反序列化json的函数根据data将数据反序列化到传入的对象中

*/

package main

import "encoding/json"
import "fmt"
import "os"

//我们用这俩个结构体来演示JSON的序列化和反序列化
type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	//我们先来看一下把基本数据类型转换成json格式。
	//布尔类型
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	//整数
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	//浮点数
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	//字符串
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	//分片，转换成json array
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	//字典，转换成json object
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	//json也能处理自定义类型，默认只转换那些可输出的字段，默认属性名为json的key。
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	//可以利用json的标签来自定义json的key，比如Page int `json:"page"`
	res2D := &Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//反序列化json字符串到go对象。
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	//使用map[string]interface{}来保存反序列化后的数据
	var dat map[string]interface{}

	//反序列化
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat)

	//对map里面的值，需要类型转换到真正的go类型。
	num := dat["num"].(float64)
	fmt.Println(num)

	//嵌套的值需要嵌套的类型转换。
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	//反序列化json数据到自定义对象，可以省去类型转换工作
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	//可以直接序列化输出到终端
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

/*

输出结果：
true
1
2.34
"gopher"
["apple","peach","pear"]
{"apple":5,"lettuce":7}
{"Page":1,"Fruits":["apple","peach","pear"]}
{"page":1,"fruits":["apple","peach","pear"]}
map[num:6.13 strs:[a b]]
6.13
a
{1 [apple peach]}
apple
{"apple":5,"lettuce":7}

*/
