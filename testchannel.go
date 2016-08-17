package main

import (
	"fmt"
	// "time"
	// "hello"
)

func g1(c chan int, t chan int) {

	fmt.Println("g1")
	for i := 0; i < 4; i++ {
		c <- int(i * 10)
	}
	close(c)
	fmt.Println("close g1")
	t <- 1
}

func g2(c chan int, t chan int) {

	for i := 0; i < 2; i++ {
		c <- i
	}
	close(c)
	fmt.Println("close g2")
	t <- 1
}

func main() {

	d1 := make(chan int)
	d2 := make(chan int)
	tag := make(chan int, 2)
	go g1(d1, tag)
	go g2(d2, tag)

	count := 0
	for {
		select {
		case gch1, ok := <-d1:
			if ok {
				fmt.Println("d1", gch1)
			}

		case gch2, ok := <-d2:
			if ok {
				fmt.Println("d2", gch2)
			}

		case i, ok := <-tag:
			if ok {
				count = i + count
				fmt.Println("index", i, count)
				if count == 2 {
					fmt.Println("Over", i)
					return
				}
			}
		}
	}

}
