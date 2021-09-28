package main

import "fmt"

//匿名字段
type person1 struct {
	string
	int
}

//嵌套结构体
type address struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	addr address
}

type company struct {
	name string
	addr address
}

//匿名嵌套结构体
type person2 struct {
	name string
	age  int
	address
}

type company2 struct {
	name string
	address
}

func main() {
	p1 := person1{"hushengs", 27}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)

	p2 := person{
		name: "hk",
		age:  27,
		addr: address{
			province: "四川",
			city:     "成都",
		},
	}

	fmt.Println(p2)
	fmt.Println(p2.addr.province)
	fmt.Println(p2.addr.city)

	p3 := person2{
		name: "hk",
		age:  27,
		address: address{
			province: "四川",
			city:     "成都",
		},
	}

	fmt.Println(p3)
	fmt.Println(p3.province)
	fmt.Println(p3.city)
}
