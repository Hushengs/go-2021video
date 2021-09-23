package main

import "fmt"

func main() {
	var s1 []int
	var s2 []string
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == nil) //true
	fmt.Println(s2 == nil) //true

	a1 := [...]int{1, 3, 5, 7, 9}
	fmt.Println(a1[1:])

	ss := make([]int, 33, 33)
	ss = append(ss, 1)
	fmt.Println(ss)
	fmt.Println(len(ss))
	fmt.Println(cap(ss))

	//本质
	s := make([]int, 0, 10)
	fmt.Printf("len(s)：%d,cap(s):%d\n", len(s), cap(s))

	//append
	s3 := []string{"北京", "上海", "广州"}
	s3 = append(s3, "深圳", "成都", "杭州", "西安")
	fmt.Println(s3)
	fmt.Printf("len(s)：%d,cap(s):%d\n", len(s3), cap(s3))

	//copy
	b1 := []int{1, 3, 5}
	b2 := b1
	b3 := make([]int, 3, 3)
	fmt.Println(b1, b2, b3)
	copy(b3, b1)
	b1[0] = 100
	fmt.Println(b1, b2, b3)

	//delete1
	x1 := [...]int{1, 3, 5}
	c1 := x1[:]
	fmt.Printf("%p\n", &c1[0])
	fmt.Printf("%v len(c1)：%d,cap(c1):%d\n", c1, len(c1), cap(c1))
	c1 = append(c1[:1], c1[2:]...)
	fmt.Printf("%p\n", &c1[0])
	fmt.Printf("%v len(c1)：%d,cap(c1):%d\n", c1, len(c1), cap(c1))
	fmt.Println(x1)
}
