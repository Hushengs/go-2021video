package main

import (
	"fmt"
	"os"
)

/*
 函数学生管理系统
*/

var (
	allstudent map[int64]*student //变量声明
)

type student struct {
	id   int64
	name string
}

func showAllStudent() {
	for k, v := range allstudent {
		fmt.Printf("学号：%d,姓名：%s\n", k, v.name)
	}
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Println("请输入学生的学号:")
	fmt.Scanln(&id)
	fmt.Println("请输入学生的姓名:")
	fmt.Scanln(&name)
	newStu := newStudent(id, name)
	allstudent[newStu.id] = newStu
}

func deleteStudent() {
	fmt.Println("请输入删除学生的学号:")
	var id int64
	var name string
	fmt.Scanln(&id)
	name = allstudent[id].name
	delete(allstudent, id)
	fmt.Printf("删除了学号为%d的学生,学生姓名为%s\n", id, name)
}

func main() {
	allstudent = make(map[int64]*student, 52)
	for {
		fmt.Println("欢迎光临学生管理系统")
		fmt.Println(`
		1、查看所有学生
		2、新增学生
		3、删除学生
		4、退出
	`)
		fmt.Println("请输入你的操作：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)

		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(-1)
		default:
			fmt.Println("滚啊！")
		}
	}
}
