package main 

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now() 

	p(now) 
	p(int(now.Month()))

	then := time.Date(2009,11,17,20,34,58,65112121,time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day()) 
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	p(then.Add(diff))
	p(then.Add(-diff))
}

// 如果需要操作时间，需要引入time包
// Year(),Day(),Hour(),Minute(),Second(),Nanosecond() 剋呀获取数字的年份和日期
// Month(),Location() ,Weekday() 获取的是字符串
// 两个时间的相减用Sub
// 两个时间的相加用Add
// 时间数据的初始化：年，月，日，时，分，秒，纳秒，时区

