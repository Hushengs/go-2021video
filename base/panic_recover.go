package main

import (
	"fmt"
)

func f1() {
	fmt.Println("1")
}

func f2() {
	//recover必须搭配defer使用
	//defer一定要在panic之前
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("释放数据库链接...")
		}

	}()
	panic("发生严重错误")
	fmt.Println("2")
}

func f3() {
	fmt.Println("3")
}

func main() {
	f1()
	f2()
	f3()
}
