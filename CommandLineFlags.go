package main 

import (
	"fmt"
	"flag"
)

func main() {
	wordPtr := flag.String("zookeeper","192.168.1.1:2181","zookeeper uri")
	numPtr := flag.Int("mappers",4,"how much mappers to be running on this task") 
	boolPtr := flag.Bool("isLocal",false,"is running on local machine") 

	var strVar string 
	flag.StringVar(&strVar,"str","google","identify the search engine")

	flag.Parse()
	fmt.Printf("zookeeper => %s , mappers => %d , isLocal => %t ,str => %s , tail => %v \n" , *wordPtr,*numPtr,*boolPtr,strVar,flag.Args()) 
}

// 要使用命令行解析工具，首先需要引入flag包
// 定义自定义变量有两种形式,下面以string形式为例： 
//  String func(name string, value string, usage string) *string
//  StringVar func(p *string, name string, value string, usage string)
//  第一种方式写起来略简洁，bool型变量默认就是false的，所以使用-isLocal不加参数就能编程true，其次--isLocal=true也是可以的
// flag.Args()可以返回出现了但为解析的命令行参数