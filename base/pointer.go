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

}
