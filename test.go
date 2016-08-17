package main

import (
	//	"fmt"
	//	"os"
	"fmt"
)

func main() {

	//	_, err := os.Create("./output.txt")
	//	a,err := os.Create("./output2.txt")
	//	if a  == nil && err !=nil {
	//
	//	}

	var b = byte(25)
	var c = &b
	var d *byte
	d = c
	var e *byte
	e = d
	fmt.Println(b, c, d, e)

}
