package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// goroutine 调度模型 GMP
	//M:N 把m个goroutine分配给n个操作系统线程去执行
	runtime.GOMAXPROCS(2) //指定多少个CPU处理，默认无需处理，默认跑满整个CPU
	//example1
	for i := 1; i < 1000; i++ {
		// go hello(i)
		go func(i int) { //传参，保证用的函数参数的i,不是外面的i
			fmt.Println(i, "hello")
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)

	//waitGroup
	f()
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go f1(j)
	}
	wg.Wait()
}

func hello(i int) {
	fmt.Println("hello", i)
}

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(10)))
	fmt.Println(i)
}
