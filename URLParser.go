package  main 

import (
	"fmt"
	"net"
	"net/url"
)

func main()  {
	s := "mysql://zmy:123456@192.168.0.1:3306/cctv?topic=kylin#20190420"
	u , _ := url.Parse(s) 
	fmt.Println(u.User)
	fmt.Println(u.User.Username())

	password , _ := u.User.Password()
	fmt.Println( password ) 
	
	host, port , _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Fragment)
	fmt.Println(u.Path)

	fmt.Println(u.RawQuery)
	m , _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["topic"])
	fmt.Println(m["topic"][0])
}


// 处理url的时候需要 net 包， net/url 包
// 解析一个url需要使用 url.Parse
// 一个url 由scheme ， username ，password ， path， query ，fragment等组成，go 分别有对应的处理函数
// 分割用户名及密码需要使用的是： net.SplitHostPort()  
// 获得query时，使用 url.ParseQuery()
