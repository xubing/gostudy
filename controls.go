package main

import (
	"fmt"
)

func TestSwitch(name string) {
	switch name {
	case "1":
		fmt.Println(name + " yes")
	case "2":
		fmt.Println(name + " NO")
	default:
		fmt.Println("default")
	}
}

func TestSwitch2(age int) int {

	switch {
	case age < 10:
		fmt.Println(string(age+'a') + " yes")
	case age < 20:
		fmt.Println(string(age+'a') + " NO")

	default:
		fmt.Println("default")
	}

	return 0

}

func main() {
	TestSwitch("1")
	TestSwitch2(2)
	TestSwitch2(11)
	TestSwitch2(21)

}
