package main

import "fmt"

func main() {
	//&取地址，*根据指针(内存地址)取值
	n := 10
	fmt.Println(&n)
	p := &n
	fmt.Printf("%T\n", p)
	m := *p
	fmt.Println(m)

	//new函数申请一个内存地址，返回对应类型指针
	var a *int
	fmt.Println(a)
	var a1 = new(int)
	fmt.Println(a1)
	fmt.Println(*a1)
	*a1 = 100
	fmt.Println(*a1)

	//make给map、slice、chan申请内存，返回对应类型本身

}
