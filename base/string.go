package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := `D:\develop\person`
	s2 := `\go2021video\理想`
	s := fmt.Sprintf("%s%s", s1, s2)
	fmt.Println(s)
	ret := strings.Split(s, "\\")
	fmt.Println(ret)
	fmt.Println(strings.Contains(s, "理想"))
	fmt.Println(strings.Contains(s, "信心"))

	fmt.Println(strings.HasPrefix(s, "理想"))
	fmt.Println(strings.HasSuffix(s, "理想"))

	fmt.Println(strings.Index(s, "理想"))

	fmt.Println(strings.Join(ret, "+"))

	//字符串修改
	s3 := "白萝卜"
	s4 := []rune(s3)
	s4[0] = '红'
	fmt.Println(string(s4))

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := "H"
	c4 := 'H'
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
}
