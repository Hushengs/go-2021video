package main

import "fmt"

type car interface {
	run()
}

type falali struct {
	brand string
}

type baoshijie struct {
	brand string
}

func drive(c car) {
	c.run()
}

func (f falali) run() {
	fmt.Printf("%s速度70迈~\n", f.brand)
}

func (b baoshijie) run() {
	fmt.Printf("%s速度70迈~\n", b.brand)
}

func main() {

	f1 := falali{
		brand: "法拉利",
	}
	b1 := baoshijie{
		brand: "保时捷",
	}

	drive(f1)
	drive(b1)

}
