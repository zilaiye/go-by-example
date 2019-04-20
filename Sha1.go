package main 

import (
	"fmt"
	"crypto/sha1"
)

func main()  {
	s := "it's a nice day today !"
	
	h := sha1.New()
	h.Write([]byte(s))
	ds := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("digest : %x \n", ds )
}
// go 如果要使用sha1，就需要引用 crypto/sha1 这个包
// 使用 sha1.New()  创建 sha1的对象
// 使用Write函数向对象中写入二进制源文
// Sum 函数返回最终的摘要
// 用到的函数列表：
// New func() hash.Hash
// Write func(p []byte) (n int, err error)
// Sum func(data []byte) [Size]byte
