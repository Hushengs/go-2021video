package main

import "fmt"

type person struct {
	name   string
	gender string
}

func f1(x person) {
	x.gender = "female" //修改的是copy的副本
}

func f2(x *person) {
	// (*x).gender = "female"
	x.gender = "female" //语法糖
}

func main() {

	var p person
	p.name = "hushengs"
	p.gender = "male"
	fmt.Println(p)
	fmt.Printf("value:%+v\n", p)
	f1(p)
	fmt.Println(p)
	fmt.Printf("value:%+v\n", p)
	f2(&p)
	fmt.Println(p)
	fmt.Printf("value:%+v\n", p)

	//获取结构体类型指针
	var p1 = new(person)
	p1.name = "conan"
	fmt.Printf("%T,%+v\n", p1, p1)
	fmt.Printf("%p\n", p1)  //person的指针，pp的值
	fmt.Printf("%p\n", &p1) //pp的内存地址

	//key-value初始化
	var p2 = &person{
		name:   "hu",
		gender: "male",
	}
	fmt.Printf("value%+v\n", p2)

	//简短初始化，顺序
	var p3 = &person{
		"ge",
		"male",
	}
	fmt.Printf("value%+v\n", p3)
	fmt.Printf("%p\n", &p3.name)
	fmt.Println(&p3.name)
	fmt.Printf("%p\n", &p3.gender)
	fmt.Println(&p3.gender)

}
