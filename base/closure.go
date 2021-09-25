package main

import "fmt"

//闭包
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func f3(x, y int) {
	fmt.Println("this is f3")
	fmt.Println(x + y)
}

//定义f4对f3进行封装
func f4(f func(int, int), x, y int) func() {
	return func() {
		// fmt.Println(x + y)
		f(x, y)
	}
}

//闭包是一个函数，这个函数包含了他外部作用域的一个变量

//底层原理：
//1、函数可以作为返回值
//2、函数内部查找变量的顺序，先在自己内部找，找不到往外层找

func main() {

	f := func(x, y int) {
		fmt.Println(x + y)
	}

	f(10, 20)

	//如果只调用一次
	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)

	ff := f2(100)
	ret := ff(120)
	fmt.Println(ret)

	f1(f4(f3, 110, 220))
}
