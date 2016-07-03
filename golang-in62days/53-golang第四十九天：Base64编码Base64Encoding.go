/*

golang第四十九天:go语言中的Base64编码Base64 Encoding
http://www.kuaizh.com/?p=761


go语言对Base64编码提供了内置的支持。

*/

package main

//这个语法引入了 encoding/base64 包并使用名称 b64代替默认的 base64。这样可以节省点空间。
import b64 "encoding/base64"
import "fmt"

func main() {

	//这是将要编解码的字符串。
	data := "abc123!?$*&()'-=@~"
	//Go 同时支持标准的和 URL 兼容的 base64 格式。编码需要使用 []byte 类型的参数，所以要将字符串转成此类型。
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	//解码可能会返回错误，如果不确定输入信息格式是否正确，那么，你就需要进行错误检查了。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	//使用 URL 兼容的 base64 格式进行编解码。
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

/*

//标准 base64 编码和 URL 兼容 base64 编码的编码字符串存在稍许不同（后缀为 + 和 -），但是两者都可以正确解码为原始字符串。

运行结果如下：
func  go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~



YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~

*/
