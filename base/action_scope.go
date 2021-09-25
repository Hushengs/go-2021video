package main

import "fmt"

var n = 100

func f1() {
	fmt.Println(n)
}

func f2() int {
	return 507
}

//函数作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func f4(a, b int) int {
	return a + b
}

//函数作为返回值
func f5(x func() int) func(int, int) int {
	// ret := func(a, b int) int {
	// 	return a + b
	// }
	// return ret
	return f4
}

func main() {
	f1()
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)

	f3(f2)
	f3(b)

	f7 := f5(f2)

	fmt.Printf("%T\n", f7)
	fmt.Println(f7)

}
