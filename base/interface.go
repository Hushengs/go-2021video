package main

import (
	"fmt"
)

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int
}

type chicken struct {
	feet int
}

func (c cat) move() {
	fmt.Println("猫动")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func (c chicken) move() {
	fmt.Println("鸡动")
}

func (c chicken) eat(food string) {
	fmt.Printf("鸡吃%s\n", food)
}

func main() {

	var a1 animal

	bluecat := cat{
		name: "mimi",
		feet: 4,
	}

	a1 = bluecat
	a1.eat("小黄鱼")
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)

	c := chicken{
		feet: 2,
	}

	a1 = c
	a1.eat("鸡饲料")
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)
}
