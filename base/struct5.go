package main

import "fmt"

//自定义类型添加方法
type myInt int

func (m myInt) hello() {
	fmt.Println("hello,我是int")
}

func main() {
	//强制类型转换
	// m := myInt(100)
	// var m = new(myInt)
	var m myInt
	m.hello()
}
