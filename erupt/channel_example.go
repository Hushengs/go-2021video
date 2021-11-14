package main

import (
	"fmt"
	"sync"
)

var wgg sync.WaitGroup
var once sync.Once

func main() {
	a := make(chan int, 50)
	b := make(chan int, 100)
	wgg.Add(3)
	go fe1(a)
	go fe2(a, b)
	go fe2(a, b)
	wgg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}

}

func fe1(ch1 chan int) {
	defer wgg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func fe2(ch1, ch2 chan int) {
	defer wgg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	once.Do(func() { close(ch2) })
}
