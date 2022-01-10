package main

import (
	"fmt"
	"go-2021video/unit_test/split"
)

func main() {
	fmt.Println("babcbedfgbc")
	ret1 := split.SplitV1("babcbedfgbc", "b")
	fmt.Printf("%#v\n", ret1)
	ret2 := split.SplitV2("babcbedfgbc", "bc")
	fmt.Printf("%#v\n", ret2)
}
