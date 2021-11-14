package main

import (
	"fmt"
	"strconv"
)

func main() {
	//数字转字符串
	i := int32(2316)
	ret := string(i) //utf8到对应字符编码 XXX错误
	fmt.Println(ret)

	ret1 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v,%T\n", ret1, ret1)
	ret11 := strconv.Itoa(2316)
	fmt.Printf("%#v,%T\n", ret11, ret11)

	//字符串转数字
	str := "1000"
	ret2, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("ParseInt Error,err:", err)
		return
	}
	fmt.Printf("%#v,%T\n", ret2, ret2)
	ret3, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Atoi Error,err:", err)
		return
	}
	fmt.Printf("%#v,%T\n", ret3, ret3)

	//字符串转bool
	boolStr := "true"
	boolVlaue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v,%T\n", boolVlaue, boolVlaue)

	//字符串转float
	floatStr := "10.24"
	floatVlaue, _ := strconv.ParseBool(floatStr)
	fmt.Printf("%#v,%T\n", floatVlaue, floatVlaue)

}
