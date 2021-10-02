package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFromFile() {
	fileObj, err := os.Open("../statics/file_read.txt")
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer fileObj.Close()
	fmt.Printf("%T", fileObj)

	// var tem := make([]byte,128)
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:]))
		if n < len(tmp) {
			return
		}
		if err == io.EOF {
			fmt.Println("read from file over")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v", err)
			return
		}
	}
}

func readFromBufio() {
	fileObj, err := os.Open("../statics/file_read.txt")
	if err != nil {
		fmt.Printf("open file failed,err:%v", err)
		return
	}
	defer fileObj.Close()
	fmt.Printf("%T", fileObj)
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		fmt.Println(line)
		if err == io.EOF {
			fmt.Println("read from file over")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v", err)
			return
		}
	}
}

func readFormFileByIoutil() {
	ret, err := ioutil.ReadFile("../statics/file_read.txt")
	if err != nil {
		fmt.Printf("read file failed,err:%v", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	// readFromFile()
	// readFromBufio()
	readFormFileByIoutil()
}
