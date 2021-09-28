package main

import "fmt"

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d dog) wang() {
	fmt.Printf("%s在叫：汪汪汪\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{name: "阿拉斯加"},
		feet:   4,
	}
	d1.wang()
	d1.move()
}
