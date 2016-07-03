/*

golang第五十天:go语言中的文件操作之读取文件Reading Files
http://www.kuaizh.com/?p=763


读写文件是最基本的功能。go语言读文件挺有意思，由于go语言的interface，使得go语言与其他语言有所不同。
与其他语言一样，go语言有File类型的结构体，但File只提供了最基本的Read，Write等功能，而类似 与ReadLine这样的功能实在bufio包里提供的。

*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//读取文件需要经常进行错误检查，这个帮助方法可以精简下面的错误检查过程。

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//也许大部分基本的文件读取任务是将文件内容读取到内存中。
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	//你经常会想对于一个文件是怎么读并且读取到哪一部分进行更多的控制。对于这个任务，从使用 os.Open打开一个文件获取一个 os.File 值开始。
	f, err := os.Open("/tmp/dat")
	check(err)

	//从文件开始位置读取一些字节。这里最多读取 5 个字节，并且这也是我们实际读取的字节数。
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	//你也可以 Seek 到一个文件中已知的位置并从这个位置开始进行读取。
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	//io 包提供了一些可以帮助我们进行文件读取的函数。例如，上面的读取可以使用 ReadAtLeast 得到一个更健壮的实现。
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//没有内置的回转支持，但是使用 Seek(0, 0) 实现。
	_, err = f.Seek(0, 0)
	check(err)

	//bufio 包实现了带缓冲的读取，这不仅对有很多小的读取操作的能提升性能，也提供了很多附加的读取函数。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	//任务结束后要关闭这个文件（通常这个操作应该在 Open操作后立即使用 defer 来完成）。
	f.Close()

}

/*

func  echo "hello" > /tmp/dat
func  echo "go" >>   /tmp/dat
func  go run reading-files.go
hello
go
5 bytes: hello
2 bytes @ 6: go
2 bytes @ 6: go
5 bytes: hello

bufio提供了很多操作，例如ReadString，ReadBytes，ReadSlice，ReadLine，使用ReadSlice和ReadLine需要小心：

1、ReadSlice很ReadLine返回的[]byte并非copy的一份副本，因此，下一次ReadSlice时，这个值就变了
2、ReadLine除了上面的问题外，ReadLine返回的数据不包括回车符\n和换行符\r

其实，ReadSting和ReadBytes已经非常好用，这两个方法的参数都是分隔符，当读到分隔符时，函数就返回。

*/
