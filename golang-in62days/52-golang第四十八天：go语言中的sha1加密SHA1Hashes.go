/*

golang第四十八天:go语言中的sha1加密SHA1 Hashes
http://www.kuaizh.com/?p=759


SHA-1是一种数据加密算法，该算法的思想是接收一段明文，然后以一种不可逆的方式将它转换成一段（通常更小）密文，
也可以简单的理解为取一串输入码（称为预映射或信息），并把它们转化为长度较短、位数固定的输出序列即散列值（也称为信息摘要或信息认证代码）的过程。

*/

package main

//Go 在多个 crypto/* 包中实现了一系列散列函数。
import "crypto/sha1"
import "fmt"

func main() {
	s := "sha1 this string"
	//产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
	h := sha1.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(s))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	//SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

/*

运行程序计算散列值并以可读 16 进制格式输出。
func  go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94

*/
