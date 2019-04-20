package main 


import (
	"fmt"
	"strconv"
)

func main() {
	i , _ := strconv.ParseInt("121211",0,32)
	fmt.Println(i) 
	j , _ := strconv.ParseInt("121211",0,16)
	fmt.Println(j) 
	k , _ := strconv.ParseUint("121211",0,16)
	fmt.Println(k) 
	l , _ := strconv.ParseInt("121211",10,64)
	fmt.Println(l)
	o , _ := strconv.ParseInt("01111111",2,8)
	fmt.Println(o)

	a , _ := strconv.ParseFloat("123.456",8) 
	fmt.Println(a)

	b , _ := strconv.ParseFloat("123.456121212121121",8) 
	fmt.Println(b)

	c , _ := strconv.ParseInt("0xabcdef",0,32)
	fmt.Println(c) 

	d , _ := strconv.ParseInt("07777", 0 , 16 ) 
	fmt.Println(d) 

	e , _ := strconv.Atoi("123456") 
	fmt.Println(e) 

	f , err  := strconv.Atoi("789.123")
	if err != nil {
		panic(err)
	}
	fmt.Println(f) 
	

}

// 将字符串的数字解析成数字
// 将字符串解析成数字需要使用的包是strconv包
// 将字符串解析成数字的函数比较多，具体选用哪个函数取决于要转换到的数字是什么类型的
// ParseInt ParseFloat  ParseUint  Atoi 
// ParseInt 的第二个参数指定了数据的机制信息，如果使用0，则会自动推测，第三个参数设置了数据存成的变量位数



