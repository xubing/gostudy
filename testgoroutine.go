package main

import (
	"fmt"
	// "time"
	// "hello"
)

func PrintB(max int) {

	for i := 0; i < max; i++ {
		defer fmt.Println(i)
		// time.Sleep(1000 * time.Millisecond) //注释掉
		fmt.Println(i + 10)

	}

}

func main() {
	// time.Sleep(1000 * time.Millisecond)
	for {
		go PrintB(2)
		// PrintB(4)
	}
	for {

	}
}
