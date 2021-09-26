package main

import "fmt"

//函数的定义
//函数是一段代码的封装
func sum(x, y int) (ret int) {
	ret = x + y
	return
}

//可变长参数
func f(x int, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

//递归
func f1(n int) int {
	if n <= 1 {
		return 1
	}
	return n * f1(n-1)
}

func main() {
	fmt.Println(sum(1, 2))
	f(1, 2, 3, 4)
	fmt.Println(f1(5))
}
