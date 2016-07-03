/*

golang第二十天:go语言中的接口Interfaces
http://www.kuaizh.com/?p=598


简单的说，interface是一组method的组合，我们通过interface来定义对象的一组行为。

interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。

与其它面向对象语言不同的是，go中无需显示声明实现了哪个接口。

*/

package main

import "fmt"
import "math"

//接口定义
type geometry interface {
	area() float64
	perim() float64
}

//在rect、circle结构体上实现接口方法
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

//rect实现了接口的所有方法

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

//circle实现了接口的所有方法

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//接口类型作为函数的参数

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	//r，c都实现了接口的所有方法，所以r和c的实例都可以作为接口的实例。
	measure(r)
	measure(c)
}

/*

func  go run interfaces.go
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
空接口（empty interface）
空接口比较特殊，它不包含任何方法：
interface{}
在Go语言中，所有其它数据类型都实现了空接口。它相当于java的Object类，是所有类的超类。

*/
