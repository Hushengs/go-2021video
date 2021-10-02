package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int63n(899999999999999999) + 100000000000000000
	fmt.Println(randNum)
}
