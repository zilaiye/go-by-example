package main 

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	fmt.Println()

	sDec , _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	fmt.Println()

	uDec , _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
	fmt.Println()

}

// 使用base64的时候需要使用到 encoding/base64 这个包
// 使用 import  别名  包名  可以给包取个别名
// base64包中有  StdEncoding 和  URLEncoding 基本的包
// base64包有StdEncoding.EncodeToString 和 StdEncoding.DecodeString 方法
// base64包还有 URLEncoding.EncodeToString() 和 URLEncoding.DecodeString() 方法
