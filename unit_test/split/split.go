package split

import "strings"

//SplitV1 切割字符串
func SplitV1(str, sep string) []string {
	var ss []string
	var tem string
	for _, v := range []rune(str) {
		if string(v) == sep {
			ss = append(ss, tem)
			tem = ""
			continue
		}
		tem = tem + string(v)
	}
	ss = append(ss, tem)
	return ss
}

func SplitV2(str, sep string) []string {
	// var ss []string
	var ss = make([]string, 0, strings.Count(str, sep)+1) //优化1
	index := strings.Index(str, sep)
	for index >= 0 {
		ss = append(ss, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	ss = append(ss, str)
	return ss
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
