package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m2 = sync.Map{} //支持并发

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n) //使用内置得Store方法设置键值对
			value, _ := m2.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
