/*

golang第十九天:go语言中的结构体方法Methods
http://www.kuaizh.com/?p=596


method是附属在一个给定的类型上的，他的语法和函数的声明语法几乎一样，只是在func后面增加了一个receiver(也就是method所依从的主体)。

在使用method的时候重要注意几点

1.虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
2.method里面可以访问接收者的字段
3.调用method通过.访问，就像struct里面访问字段一样
struct的method的形式如下：

    func (r ReceiverType) funcName(parameters) (results)

如果想要修改struct的成员的值，method被定义时候其ReceiverType必须是struct*形式。如果ReceiverType是struct，则无法改变struct成员的值。

*/

package main

import "fmt"

type rect struct {
	width, height int
}

//接受者为结构体指针
func (r *rect) area() int {
	return r.width * r.height
}

//接受者为结构体
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
func main() {
	r := rect{width: 10, height: 5}
	//调用
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
	//go自动处理指针和结构体之间的转换。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

/*

$ go run methods.go
area:  50
perim: 30
area:  50
perim: 30


*/
