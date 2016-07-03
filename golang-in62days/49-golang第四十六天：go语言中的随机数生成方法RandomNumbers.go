/*

golang第四十六天:go语言中的随机数生成方法Random Numbers
http://www.kuaizh.com/?p=753


golang生成随机数可以使用math/rand包


*/

package main

import "time"
import "fmt"
import "math/rand"

func main() {

	////例如，rand.Intn 返回一个随机的整数 n，0 <= n <= 100。
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	//rand.Float64 返回一个64位浮点数 f，0.0 <= f <= 1.0。
	fmt.Println(rand.Float64())

	//这个技巧可以用来生成其他范围的随机浮点数，例如5.0 <= f <= 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	//要让伪随机数生成器有确定性，可以给它一个明确的种子。
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	//调用上面返回的 rand.Source 的函数和调用 rand 包中函数是相同的。
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	//如果使用相同的种子生成的随机数生成器，将会产生相同的随机数序列。
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}

/*

$ go run random-numbers.go
81,87
0.6645600532184904
7.123187485356329,8.434115364335547
0,28
5,87
5,87

*/
