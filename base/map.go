package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var m1 map[string]int
	fmt.Println(m1)
	m1 = make(map[string]int, 10) //尽量估算好map容量
	m1["hushengs"] = 28
	m1["conan"] = 13
	fmt.Printf("%T\n", m1)
	fmt.Printf("%v", m1)
	fmt.Println(m1["super"])
	if value, ok := m1["super"]; !ok {
		fmt.Println("不存在对应key")
	} else {
		fmt.Println(value)
	}
	//delete
	delete(m1, "conan")
	fmt.Println(m1)
	delete(m1, "super")
	fmt.Println(m1)

	//排序
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)
	fmt.Println(len(scoreMap))
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("student%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)

	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	//<--元素类型为map的切片
	// var sm1 = make([]map[int]string, 0, 10)
	// sm1[0][100] = "A"
	// fmt.Println(sm1)
	// // index out of range

	// var sm2 = make([]map[int]string, 1, 10)
	// sm2[0][100] = "A"
	// fmt.Println(sm2)
	// //entry in nil map

	//slice与map都要初始化
	var sm1 = make([]map[int]string, 1, 10)
	sm1[0] = make(map[int]string, 0)
	sm1[0][100] = "A"
	fmt.Println(sm1)

	//<--值为切片类型的map
	var ms1 = make(map[string][]int, 10)
	ms1["hushengs"] = []int{10, 20, 30}
	fmt.Println(ms1)
}
