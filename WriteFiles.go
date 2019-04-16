package main 

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("nice\nto\tmeet you\n")
	err := ioutil.WriteFile("/tmp/dat1",d1,0644)
	check(err)

	f ,err := os.Create("/tmp/dat2") 
	defer f.Close()

	d2 := []byte{92,93,94}
	n1 , err := f.Write(d2)
	check(err)
	fmt.Printf("write %d bytes \n",n1 )

	n2 , err := f.WriteString("fuck you man !") 
	check(err) 
	fmt.Printf("write %d bytes \n",n2) 
	
	f.Sync()

	w := bufio.NewWriter(f)
	n3 , err := w.WriteString("buffered !\n") 
	check(err) 
	fmt.Printf("write %d bytes !\n",n3)
	w.Flush()

}


// 技术要点：
// 1. go的写入文件可以使用io/ioutil包，ioutil.WriteFile可以直接写入二进制数据到文件中
// 2. os 包可以创建文件，os.Create(filename string),创建文件后一定要记得关闭，defer 关键字可以帮助干这个事
// 3. 打开的文件可以写入[]byte,使用Write函数，也可以直接写入字符串，WriteString
// 4. 打开的文件同样可以用来创建带缓存的Writer,bufio.NewWriter()可以创建一个带缓存的writer对象，这个对象有WriteString方法
// 函数列表：
// 1. WriteFile func(filename string, data []byte, perm os.FileMode) error
// 2. Create(filename string) error 
// 3. Write func(b []byte) (n int, err error)
// 4. WriteString func(str string) (n int ,err error) 
// 5. Close()  Sync() 
// 6. NewWriter func(w io.Writer) *Writer
// 7. WriteString func(s string) (int, error)
// 8. Flush()
