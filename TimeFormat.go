package main 

import (
	"fmt"
	"time"
)

func main()  {
	p := fmt.Println
	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, _ := time.Parse(time.RFC3339, "2019-04-01T11:12:13+01:02")
	p(t1) 
	
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("Mon Jan 2 15:04:05 2006"))
	p(t.Format("2006/01/02T15:04:05.999999-07:00"))

	form := "3 04 PM"
	t2 , _ := time.Parse(form , "8 41 AM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", 
			t.Year(),t.Month(),t.Day(),
			t.Hour(),t.Minute(),t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_ ,e := time.Parse(ansic, "10:20PM") 
	p(e) 
}

// go 中使用时间函数需要调用time包
// time.Now() 获取当前的时间
// 对于一个时间可以调用Format函数，并辅助特定的格式
// time.Parse(fmt,string) 可以根据格式分析字符串
// go的时间格式化函数比较特别，是根据go的诞生时间来定的，并非java中的%Y%m%d 之类的
// go的诞生日期是 ： 2006-01-02T15:04:05.999999-07:00 
// 2019-04-21T21:51:41+08:00 是一个标准的RFC3339格式的时间
// Year(), Month(),Day(),Hour(),Minute(),Second()
// 变态的go非要你记住它的诞生日期作为时间格式。。。。。。

