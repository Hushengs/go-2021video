package main

import (
	"fmt"
)

//iota go语言常量计数器
//const关键字出现是iota被置为零
//每新增一行常量声明，iota加1
const (
	a1 = iota
	a2
	a3
	a4
)

func main() {
	fmt.Println(a1, a2, a3, a4)
}
