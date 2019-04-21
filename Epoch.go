package main 

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	sec := now.Unix()
	nano := now.UnixNano()
	fmt.Println(now)
	fmt.Println(sec)
	fmt.Println(nano)

	millis := nano/10e6 
	fmt.Println(millis)

	fmt.Println(time.Unix(sec,0))
	fmt.Println(time.Unix(0,nano))
	fmt.Println(time.Unix(sec,nano)) // this is a bad example 
}

// go 中调用时间需要使用time包
// go中的获取某个时间对应的秒使用Unix()函数
// go中获取某个时间对应的纳秒使用UnixNano()函数
// 将纳秒除以1000000转换成微秒 
// time.Unix(秒，纳秒) 既可以接受纳秒，也可以接受微秒 , 如果提供了两个会累加
