package main 

import (
	"os"
	"fmt"
)

func main() {
	argsWithProg := os.Args 
	argsWithoutProg := os.Args[1:]
	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

// 这个demo确实比较简陋，类似于c，c++， go有一个收集参数的变量os.Args，我们看到go的main函数是没有入参的，flag可能是更好的获取方法
