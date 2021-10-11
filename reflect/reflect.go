package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	fmt.Printf("type name:%v ;type kind:%v\n", v.Name(), v.Kind())
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	str := `{"name":"Hushengs","age":27}`
	var p person
	json.Unmarshal([]byte(str), &p)
	fmt.Printf("%+v\n", p)

	pt := reflect.TypeOf(p)
	pv := reflect.ValueOf(p)
	fmt.Println(pt, pv)

	var a float32
	reflectType(a)

	var pp person
	reflectType(pp)

	var i int64 = 100
	// reflectSetValue(i)
	reflectSetValue(&i)
	fmt.Println(i)
	reflectSetValue1(&i)
	fmt.Println(i)


}
