package main

import "fmt"

func main() {
	fmt.Print("Hushengs")
	fmt.Print(" GO FOR IT")

	fmt.Println("hushengs")
	fmt.Println("go for it")

	var n = 100
	var s = "hushengs"
	var b = false
	//查看类型
	fmt.Printf("%T\n", n)
	//二进制
	fmt.Printf("%b\n", n)
	//八进制
	fmt.Printf("%o\n", n)
	//十进制
	fmt.Printf("%d\n", n)
	//十六进制
	fmt.Printf("%x\n", n)
	//字符
	fmt.Printf("%c\n", []byte(s)[0])
	//字符串
	fmt.Printf("%s\n", s)
	//指针
	fmt.Printf("%p\n", &n)
	//值
	fmt.Printf("%v\n", n)
	//浮点数
	fmt.Printf("%f\n", float64(n))
	fmt.Printf("%f\n", float32(n))
	//布尔值
	fmt.Printf("%t\n", b)

	type person struct {
		name string
	}
	people := new(person)
	people.name = "conan"
	fmt.Printf("%v\n", people)
	//+v 输出结构体时会添加字段名
	fmt.Printf("%+v\n", people)
	//#v go语法表示
	fmt.Printf("%+v\n", people)
	//值类型
	fmt.Printf("%T\n", people)
	//百分号
	fmt.Printf("%d%%\n", n)

	//整形
	fmt.Println("<----------------------整形---------------------------->")
	//二进制
	fmt.Printf("%b\n", n)
	//八进制
	fmt.Printf("%o\n", n)
	//十进制
	fmt.Printf("%d\n", n)
	//十六进制a-f
	fmt.Printf("%x\n", n)
	//字符
	fmt.Printf("%c\n", n)
	//十六进制a-f
	fmt.Printf("%X\n", n)
	//Unicode格式
	fmt.Printf("%U\n", n)
	//该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	fmt.Printf("%q\n", n)

	fmt.Println("<----------------------宽度符---------------------------->")
	var f = 12.34
	//默认宽度默认精度
	fmt.Printf("%f\n", f)
	//宽度9,默认精度
	fmt.Printf("%9f\n", f)
	//默认宽度,精度2
	fmt.Printf("%.2f\n", f)
	//宽度9，精度2
	fmt.Printf("%9.2f\n", f)
	//宽度9，精度
	fmt.Printf("%9.f\n", f)

	//获取用户输入
	// var ss string
	// fmt.Scan(&ss)
	// fmt.Println("用户输入的内容是：", ss)

	//指定结束符，多个值
	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)
}
