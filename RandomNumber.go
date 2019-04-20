package main 

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println(rand.Intn(100)) 
	for i := 0 ; i < 20 ; i++ {
		fmt.Print(rand.Intn(100),",") 
		if i % 10 == 0 {
			fmt.Println()
		}
	}
	
	s1 := rand.NewSource(time.Now().UnixNano()) 
	r1 := rand.New(s1)
	for j := 0 ; j < 20 ; j++ {
		fmt.Print(r1.Intn(20),",")
	}
	fmt.Println()

	for l := 0 ; l < 10 ; l++ {
		fmt.Print(r1.Float64(),",")
	}
	fmt.Println()

	s2 := rand.NewSource(11) 
	r2 := rand.New(s2) 
	for k := 0 ; k < 20 ; k++ {
		fmt.Print(r2.Intn(20),",")
	}
	fmt.Println()

	fmt.Println(rand.Float64()*5+2)
	fmt.Println(rand.Float64()*10 )


}

// 随机数的调用需要使用到 math/random 库
// time 库可以给随机数生成种子
// rand.Intn(100) 可以生成100以内的随机数
// rand.Float64() 可以生成浮点数,这个浮点数在0到1之间
// rand.NewSource(seed) 可以设置种子
// time.Now().UnixNano() 可以获取unix时间戳，unix时间戳可以作为种子
