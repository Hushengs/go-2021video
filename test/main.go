package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileObj, err := os.OpenFile("../statics/logfile.txt", os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	log.SetOutput(fileObj)
	for {
		log.Println("这是测试日志")
		time.Sleep(3 * time.Second)
	}
}
