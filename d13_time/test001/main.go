package main

import (
	"fmt"
	"time"
)

func test0() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Unix())
	fmt.Println(now.Date())
	fmt.Println(now.Clock())
	fmt.Printf("%T, %v\n", now.Year(), now.Year())
	fmt.Printf("%T, %v\n", now.Month(), now.Month())
	fmt.Printf("%T, %v\n", now.Day(), now.Day())
	fmt.Printf("%T, %v\n", now.Hour(), now.Hour())
	fmt.Printf("%T, %v\n", now.Minute(), now.Minute())
	fmt.Printf("%T, %v\n", now.Second(), now.Second())
}

func test1() {
	now := time.Now()
	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Printf("%T\n", timestamp1)
	fmt.Printf("%T\n", timestamp2)
}

func test2() {
	timeObj := time.Unix(1603158831, 0)
	fmt.Println(timeObj)
}

func test3() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))

	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	fmt.Println(now.Format("2006-01-02 15:04:05 Mon Jan"))
}

func test4() {
	datetimeStr := "2020-10-20 10:40:47"

	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", datetimeStr, loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
}

func test5() {
	d := time.Saturday
	fmt.Printf("%T, %v\n", d, d == 6)
}

func test6() {
	today := time.Now()
	tomorrow := today.Add(time.Hour * 24)
	yesterday := today.Sub(time.Hour * 24)

	fmt.Println(today)
	fmt.Println(tomorrow)
	fmt.Println(yesterday)
}

func main() {
	test6()
}
