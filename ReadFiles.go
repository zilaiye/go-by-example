package main 

import (
	"fmt"
	"io/ioutil"
	"io"
	"os"
	"bufio"
)
func check(err error) {
	if err != nil {
		panic(err) 
	}
}

func main() {
	dat , err := ioutil.ReadFile("/tmp/dat1") 
	check(err) 
	fmt.Printf("read %d bytes , contents %s \n", len(dat),string(dat))

	f , err := os.Open("/tmp/dat2") 
	check(err) 
	defer f.Close()
	
	b1 := make([]byte , 5) 
	n1 , err := f.Read(b1)
	check(err) 
	fmt.Printf("read %d bytes , contents: %s \n", n1, string(b1))

	b2 := make([]byte , 100) 
	n2 , err := f.Read(b2) 
	check(err) 
	fmt.Printf("read %d bytes , content : %s \n", n2, string(b2))

	o3 , err := f.Seek(6,0) 
	check(err) 
	b3 := make([]byte,2) 
	n3 , err := io.ReadAtLeast(f,b3,2) 
	check(err) 
	fmt.Printf("%d bytes @ %d : %s \n", n3, o3, string(b3) )

	_ , err = f.Seek(0,0) 
	check(err) 

	br := bufio.NewReader(f) 
	b4,err := br.Peek(5) 
	check(err) 
	fmt.Printf("5 bytes : %s \n",string(b4)) 

}

// 1. go读写文件可能用到的包有 io，os，bufio
// 2. go 读取文件可以使用 ioutil.ReadFile,这届读取文件所有内容
// 3. go 打开文件，使用os.Open接口
// 4. go 读取指定字节的文件使用 Read函数 ，读取之前需要先申请缓存
// 5. go 可以重定义字节流的位置，找到合适的位置
// 6. 除了使用Read 还可以使用ReadAtLeast函数读取指定字节的数据
// 7. 使用bufio.NewReader 来创建带缓存的阅读器，使用Peek函数可以读取指定的字节数

// go的数据读取用到的函数有：
// ioutil.ReadFile  =>  ReadFile func(filename string) ([]byte, error)
// os.Open => Open func(name string) (*File, error)
// Read => Read func(b []byte) (n int, err error)
// io.ReadAtLeast => ReadAtLeast func(r Reader, buf []byte, min int) (n int, err error)
// bufio.NewReader => NewReader func(rd io.Reader) *Reader
// Seek => Seek func(offset int64, whence int) (ret int64, err error)
// Peek => Peek func(n int) ([]byte, error)

