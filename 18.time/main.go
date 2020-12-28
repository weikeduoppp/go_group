package main

import (
	"fmt"
	"time"
)

func _func() {
	var now time.Time // 时间对象
	now = time.Now()
	fmt.Printf("%v\n", now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func unix() {
	now := time.Now()
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
	time2 := time.Unix(timestamp1, 0) //将时间戳转为时间格式
	fmt.Println(time2.Year())
	fmt.Println(time2.YearDay())

}

func durations() {
	fmt.Println()
	fmt.Println("time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位:")
	// time.Duration表示1纳秒，time.Second表示1秒。
	const (
		Nanosecond  time.Duration = 1
		Microsecond               = 1000 * Nanosecond
		Millisecond               = 1000 * Microsecond
		Second                    = 1000 * Millisecond
		Minute                    = 60 * Second
		Hour                      = 60 * Minute
	)
	const duration time.Duration = 1
	fmt.Println(time.Second)
	fmt.Println(duration)
	fmt.Println("时间操作:")
	now := time.Now()
	// 要求时间+时间间隔的需求  func (t Time) Add(d Duration) Time
	fmt.Println(now.Add(Hour)) // 一个小时之后的时间
	/*
		Sub 两个时间之间的差值
		Equal 判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较
		Before 如果t代表的时间点在u之前，返回真；否则返回假。 func (t Time) Before(u Time) bool
		After 如果t代表的时间点在u之后，返回真；否则返回假。
	*/
}

func func2() {
	// 定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	// time对象格式化
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))

	// 将时间模板转化回time
	t, err := time.Parse("2006/01/02", "2020/12/18")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Unix())
}

func main() {
	fmt.Println()
	fmt.Println("time:")
	_func()
	fmt.Println()
	fmt.Println("时间戳:")
	unix()
	durations()
	fmt.Println()
	fmt.Println("定时器 Tick && 格式化(2006 1 2 3 4 5) Format:")
	func2()
}
