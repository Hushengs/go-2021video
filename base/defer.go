package main

import (
	"fmt"
)

//go语音中的函数的return不是原子操作，在底层是分为两步执行的
//1、返回值赋值
//2、真的RET返回
//3、如果有defer，执行时间在1,2之间

//5
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

//6
func f2()(x int){
	defer func(){
		x++
	}
	return 5
}

//5
func f3()(y int){
	x := 5
	defer func(){
		x++
	}()
	return x
}

//5
func f4()(x int){
	defer func(x int){
		x++
	}(x)
	return 5
}

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("hushengs")
	defer fmt.Println("hello")
	fmt.Println("end")
}

func main() {
	deferDemo()
}
