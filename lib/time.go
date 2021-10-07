package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	//time.Unix()
	ret := time.Unix(1633163940, 0)
	fmt.Println(ret.Year())
	fmt.Println(ret.Month())
	fmt.Println(ret.Day())
	//时间间隔
	fmt.Println(time.Second)
	//时间操作
	fmt.Println(now.Add(24 * time.Hour))

	//定时器
	// timer := time.Tick(time.Second)
	// for t := range timer {
	// 	fmt.Println(t)
	// }

	//格式化时间
	fmt.Println(now.Format("2006-01-02"))
	//UTC
	timeObj, err := time.Parse("2006-01-02", "2021-10-02")
	if err != nil {
		fmt.Printf("Parse time err%v\n", err)
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
	//时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load loc failed, err%v\n", err)
	}
	cstObj, err := time.ParseInLocation("2006-01-02", "2021-10-02", loc)
	if err != nil {
		fmt.Printf("Parse time err%v\n", err)
	}
	fmt.Println(cstObj)
	fmt.Println(cstObj.Unix())
	//差值
	d := now.Sub(timeObj)
	d1 := now.Sub(cstObj)
	fmt.Println(d)
	fmt.Println(d1)
	//sleep
	fmt.Println("开始sleep")
	time.Sleep(5 * time.Second)
	fmt.Println("5秒钟过去了")
}
