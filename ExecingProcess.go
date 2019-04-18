package main 

import (
	"syscall"
	"os"
	"os/exec"
	"fmt"
)

func main() {
	// 找到ls的绝对路径
	lsBinary , lookErr := exec.LookPath("ls") 
	if lookErr != nil {
		panic(lookErr)
	}
	fmt.Println(lsBinary)

	args := []string{"ls","-lrta"}
	env := os.Environ()
	lsErr := syscall.Exec(lsBinary,args,env)
	if lsErr != nil {
		panic(lsErr)
	}
}

// 创建子进程的包有 os/exec , syscall
// 使用exec.LookPath() 可以找到bin文件的绝对路径
// 使用 syscall.Exec可以发起子调用，需要bin文件绝对路径，参数列表及系统的环境变量
// LookPath func(file string) (string, error)
// Environ func() []string
// Exec func(argv0 string, argv []string, envv []string) (err error)

