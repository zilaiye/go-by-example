package main 

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	os.Setenv("WHOISYOURDADDY","Kylin")
	fmt.Println("WHOISYOURDADDY : " , os.Getenv("WHOISYOURDADDY"))
	fmt.Println("WHOISYOURFather :" , os.Getenv("WHOISYOURFather"))
	fmt.Println()

	for _ , e := range os.Environ() {
		// fmt.Println(e)
		pair := strings.Split( e, "=")
		fmt.Println(pair[0]," --> ", pair[1])
	}
}

// 这个demo的重点是环境变量的设置和读取问题
// 使用os.Setenv 设置环境变量
// 使用os.Getenv 获取环境变量的值
// 使用os.Environ 遍历环境变量
// Setenv func(key, value string) error
// Getenv func(key string) string
// Environ func() []string

