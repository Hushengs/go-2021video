package main

import "fmt"

//1.使用&来判断一个数字是奇数还是偶数
//3 0011
//1 0001
//位与
func oddOrEven(num int64) {
	if num&1 == 1 {
		fmt.Printf("%d is odd(奇数) \n", num)
	} else {
		fmt.Printf("%d is even(偶数数) \n", num)
	}
}

func main() {
	oddOrEven(111)
	oddOrEven(112)
}
