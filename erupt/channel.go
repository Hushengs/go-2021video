package main

import (
	"fmt"
	"sync"
)

var c chan int
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	fmt.Println(c)
	c = make(chan int) //不带缓冲区的通道初始化
	go func() {
		defer wg.Done()
		ret := <-c
		fmt.Println("从通道中取到值", ret)
	}()
	c <- 1
	fmt.Println("1发送到通道中")
	cc := make(chan int, 16) //带缓冲区的通道初始化
	fmt.Println(c)
	fmt.Println(cc)
	close(c)
	close(cc)
	fmt.Println(c)
	fmt.Println(cc)
	wg.Wait()
}
