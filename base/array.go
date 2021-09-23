package main

import "fmt"

func main() {
	citys := [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	//多维数组
	var all [3][2]int
	all = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(all)

	
}
