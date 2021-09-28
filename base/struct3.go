package main

import "fmt"

type person struct {
	name string
	age  int
}

//构造函数
//返回结构体还是结构体指针(当结构体大的时候使用结构体指针，减少程序的内存开销)
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("hk", 28)
	p2 := newPerson("conan", 13)
	fmt.Println(p1, p2)
}


