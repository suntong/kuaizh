/*

golang第三十八天:go语言中的资源回收功能Defer
http://www.kuaizh.com/?p=735


defer的思想类似于C++中的析构函数，不过Go语言中"析构"的不是对象，而是函数，defer就是用来添加函数结束时执行的语句。
注意这里强调的是添加，而不是指定，因为不同于C++中的析构函数是静态的，Go中的defer是动态的。
但是，要注意的是，如果我们的defer语句没有执行，那么defer的函数就不会添加。
defer用于保证那个函数在程序退出时会被执行，像其他语言的ensure或者finally一样，通常用来做清理工作。



*/

package main

import "fmt"
import "os"
	
//创建文件，写入内容，然后关闭文件。

func main() {
	//创建文件后，定义defer，然后main函数退出时会调用defer后面的语句。实际上是writeFile先调用，然后再调用closeFile
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

/*



func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
}

/*

    return f
}

/*



func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

/*



func closeFile(f *os.File) {
    fmt.Println("closing")
    f.Close()
}

/*

	

执行结果：
creating
writing
closing

*/
