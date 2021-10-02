package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func writeFile() {
	fileObj, err := os.OpenFile("../statics/file_write.txt", os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
	}
	defer fileObj.Close()
	fileObj.Write([]byte("哈哈哈1"))
	fileObj.WriteString("哈哈哈2")
	io.WriteString(fileObj, "哈哈哈3")
}

func writeFileBufio() {
	fileObj, err := os.OpenFile("../statics/file_write.txt", os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
	}
	defer fileObj.Close()
	//创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello Hushengs")
	wr.Flush()
}

func writeFileIoutil() {
	err := ioutil.WriteFile("../statics/file_write.txt", []byte("hello conan"), 0666)
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
	}
}

func main() {
	// writeFile()
	// writeFileBufio()
	writeFileIoutil()
}
