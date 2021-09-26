package main

import "fmt"

//自定义类型和类型别名
type myInt int    //自定义类型
type youInt = int //类型别名

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m youInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var p person
	p.name = "hushengs"
	p.age = 28
	p.gender = "male"
	p.hobby = []string{"游泳", "乒乓球"}
	fmt.Println(p)
	fmt.Printf("type:%T\n", p)
	fmt.Println(p.name)
	fmt.Printf("value:%+v\n", p)

	//匿名结构体
	var s struct {
		x string
		y int
	}
	s.x = "hehe"
	s.y = 100
	fmt.Printf("type:%T value:%+v\n", s, s)
}
