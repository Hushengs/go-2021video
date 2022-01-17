package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:mysql@tcp(127.0.0.1:33061)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("ping:%s failed,err:%v\n", dsn, err)
		return
	}
	fmt.Println("打开数据库成功！")
}
