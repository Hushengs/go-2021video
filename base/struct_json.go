package main

import (
	"encoding/json"
	"fmt"
)

//结构体与json

//1、序列化：把go语音中的结构体变量--》json格式字符串
//2、反序列化：json格式的字符串--》go语音中的结构体变量
// tag

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "hushegns",
		Age:  27,
	}
	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed err:%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	//反序列化
	str := `{"name":"conan","age":13}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)

}
