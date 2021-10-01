package main

import (
	"fmt"
)

//接口可以嵌套
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

//同一结构体可以实现多个接口
type cat struct {
	name string
	feet int
}

func (c cat) move() {
	fmt.Println("猫动")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func main() {

	var a1 animal

	c1 := cat{
		"假老练",
		4,
	}
	c1.move()
	a1 = c1
	a1.eat("风车车")
}
