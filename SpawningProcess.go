package main 

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)


func main() {
	dateCmd := exec.Command("date")
	date , err  := dateCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(" > date ") 
	fmt.Println(string(date))

	grepCmd :=  exec.Command("grep","hello") 

	grepIn , _ := grepCmd.StdinPipe()
	grepOut , _ := grepCmd.StdoutPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hello , nice to meet you \n god is a girl"))
	grepIn.Close()

	grepBytes , _ := ioutil.ReadAll(grepOut) 
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut , err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("ls -a -l -h ") 
	fmt.Println(string(lsOut))


}

// 创建子进程需要使用的包是 os/exec 包
// 使用exec.Command 来创建子进程对象
// 子进程对象有多个方法： 比如说 Start() , Close() , Wait() , Output() , StdinPipe() ,StdoutPipe() 
// ioutil.ReadAll() 可以读取管道的输出信息

// 函数列表
// Command func(name string, arg ...string) *Cmd
// Output func() ([]byte, error)
// StdinPipe func() (io.WriteCloser, error)
// StdoutPipe func() (io.ReadCloser, error)
// Start func() error
// Write func(p []byte) (n int, err error)
// ReadAll func(r io.Reader) ([]byte, error)
// Wait func() error



