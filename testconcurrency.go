package main

// https://divan.github.io/posts/go_concurrency_visualize/
import "time"
import "fmt"

func main() {
	var Ball int
	table := make(chan int)
	go player(table, 1)
	go player(table, 2)

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table
}

func player(table chan int, playerid int) {
	for {
		ball := <-table
		fmt.Println(playerid, ":", ball)
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
