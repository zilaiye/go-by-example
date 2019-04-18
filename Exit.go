package main 

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")

	os.Exit(3)
}

// defer 遇到 os.Exit() 是不会执行的，所以 ！永远不会输出
// go 语言的main是不会返回int结果的，go的输出结果会被go捕获
// 如果想看到推出状态3，你需要编译成exe文件执行，然后echo $?  
