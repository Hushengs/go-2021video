package main

import "fmt"

func main() {
	m1 := make(map[string]interface{}, 16)
	m1["name"] = "hushengs"
	m1["age"] = 1
	m1["merried"] = false
	m1["hobby"] = [...]string{"游泳", "跑步"}
	fmt.Println(m1)
}
