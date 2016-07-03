/*

golang第二十一天:go语言中的错误处理Errors
http://www.kuaizh.com/?p=600


在 Golang 中，错误处理机制一般是函数返回时使用的，是对外的接口，而异常处理机制 panic-recover 一般用在函数内部。

error 类型实际上是实现了Error()方法的类型。error接口定义了Error()方法，Golang 使用该接口进行标准的错误处理。

type error interface {
Error() string
}

一般情况下，如果函数需要返回错误，就将 error 作为多个返回值中的最后一个（但这并非是强制要求）。参考模型：

*/

package main

import "errors"
import "fmt"

//error作为最后一个返回值。

func f1(arg int) (int, error) {
	if arg == 42 {
		//errors.New构建一个错误
		return -1, errors.New("can’t work with 42")
	}
	//没有错误，返回nil
	return arg + 3, nil
}

//自定义错误类型：结构体实现接口的方法，结构体实现了error接口。
type argError struct {
	arg  int
	prob string
}

//实现error接口的方法
func (e *argError) Error() string {
	return fmt.Sprintf("%d – %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		//&argError生成一个错误对象
		return -1, &argError{arg, "can’t work with it"}
	}
	return arg + 3, nil
}

func main() {
	//错误处理测试
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	//如果要获取自定义的错误的内部信息，需要使用断言e.(*argError)。判断错误是否是我们定义的类型。
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

/*

输出结果：
f1 worked: 10
f1 failed: can’t work with 42
f2 worked: 10
f2 failed: 42 – can’t work with it
42
can’t work with it

*/
