package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Test map")
	var m = make([]int, 5)
	for index := 0; index < 10; index++ {
		m[index] = index
		fmt.Println(m[index])
	}
	fmt.Println(m, len(m))

	strings.EqualFold()
}
