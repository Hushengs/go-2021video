package main

import "fmt"

type cat struct{}

type dog struct{}

type person struct{}

type speaker interface {
	speak()
}

//引出接口的例子
func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func (p person) speak() {
	fmt.Println("啊啊啊")
}

func hit(x speaker) {
	//接收一个参数，传进来什么，我就打什么
	x.speak()
}
func main() {
	var c1 cat
	var d1 dog
	var p1 person
	// c1.speak()
	// d1.speak()
	// p1.speak()
	hit(c1)
	hit(d1)
	hit(p1)
}
