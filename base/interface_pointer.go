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

func (c *cat) move() {
	fmt.Println("猫动")
}

func (c *cat) eat(food string) {
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

	// c1 := cat{"tom",4} //cat
	c2 := &cat{"假老练",4}

	// a1 = c1
	// fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
