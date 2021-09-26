package main

import "fmt"

type dog struct {
	name string
	age  int
}

//构造函数
//返回结构体还是结构体指针(当结构体大的时候使用结构体指针，减少程序的内存开销)
func newDog(name string, age int) dog {
	return dog{
		name: name,
		age:  age,
	}
}

//方法是作用于特定类型的函数
//接受者表示的是调用该方法的具体类型变量，多用类型名字首字母小写 d
//值接受者，指针接受者
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪", d.name)
}

func (d *dog) guonian() {
	d.age += 1
}

func (d *dog) dream() {
	fmt.Printf("%s:梦想睡觉", d.name)
}

func main() {
	d1 := newDog("阿拉斯加", 1)
	fmt.Printf("%+v\n", d1)
	d1.wang()
	d1.guonian()
	fmt.Printf("%+v\n", d1)
	d1.dream()
}
