package main 

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	sigs := make(chan os.Signal, 1) 
	done := make(chan bool , 1 ) 

	signal.Notify(sigs,syscall.SIGINT,syscall.SIGTERM)
	go func() {
		s := <- sigs
		fmt.Println()
		fmt.Println("signals:", s ) 
		done <- true 
	}() 
	fmt.Println("awaiting...") 
	<- done 
	fmt.Println("exiting") 
}

// go的信号处理需要使用的包有： os , os/signal , syscall 
// signal有特殊的类型，那就是os.Signal
// signal.Notify 可以捕获系统的给进程发出的信号
// 函数列表：
// Notify func(c chan<- os.Signal, sig ...os.Signal)
