/*

golang第十八天:go语言中的结构体Structs
http://www.kuaizh.com/?p=594


Go语言中，也有struct，定义与C语言类似，它是字段的集合。它可以很好的把字段组织成记录。

*/

package main

import "fmt"

//person有俩个字段。
type person struct {
	name string
	age  int
}

func main() {
	//生成一个结构体实例
	fmt.Println(person{"Bob", 20})
	//初始化时指定名称
	fmt.Println(person{name: "Alice", age: 30})
	//未指定的将初始化成默认值
	fmt.Println(person{name: "Fred"})
	//指向结构体的指针
	fmt.Println(&person{name: "Ann", age: 40})
	//采用.来访问结构体的字段
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	//.运算符也可以用于结构体指针
	sp := &s
	fmt.Println(sp.age)
	//结构体是可变的
	sp.age = 51
	fmt.Println(sp.age)
}

/*

$ go run structs.go
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
Sean
50
51

Go语言支持只提供类型，而不写字段名的方式，也就是匿名字段，或称为嵌入字段。当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。下面举例说明之：

type Human struct {
name string
age int
}

type Student struct {
Human        //匿名字段，那么默认Student就包含了Human的所有字段
speciality string
}

//初始化Student
mark := Student(Human{"shicq", 31}, "Computer Science")

//访问相应字段
fmt.Println("His name is ", mark.name)
fmt.Println("His age is ", mark.age)
fmt.Println("His speciality is ", mark.speciality)

我们看到Student访问属性age和name的时候，就像访问自己所拥有的字段一样。当然Student也能通过访问Human来访问这两个字段：
mark.Human = Human{"shicq", 31}
mark.Human.age -= 1

*/
