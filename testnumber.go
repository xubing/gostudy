package main

import "fmt"

func main() {

	//overflow
	var u8 uint8 = 255
	fmt.Println(u8, u8+1, u8+2)

	var u = 255
	fmt.Println(u, u+1, u+2)

	//bit operator
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)
	fmt.Println("")
	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)
}
