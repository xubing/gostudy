package main

import "fmt"

func squares(index int) func() int {
	var x int = index
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f := squares(1)
	fmt.Println(f())

	//fmt.Println(f(3)) // "1"
	//fmt.Println(f(4)) // "4"
	//fmt.Println(f(5)) // "9"
	//fmt.Println(f()) // "16"
}
