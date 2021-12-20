package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex
var rwLock sync.RWMutex

func add() {
	for i := 0; i < 5; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func read() {
	defer wg.Done()
	lock.Lock()
	// rwLock.RLock()
	fmt.Println(x)
	time.Sleep(time.Microsecond)
	lock.Unlock()
	// rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	// rwLock.Lock()
	// x += 1
	x = x + 1
	time.Sleep(time.Microsecond * 5)
	lock.Unlock()
	// rwLock.Unlock()
}
func main() {
	// wg.Add(2)
	// go add()
	// go add()
	// wg.Wait()
	// fmt.Println(x)
	start := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
