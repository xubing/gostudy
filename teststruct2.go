package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Vector struct {
	X int
	Y int
}

type Component struct {
	Point
	Vector
}

func main() {
	var t1 = Component{Point{1, 2}, Vector{3, 4}}
	fmt.Println(t1, t1.Point.X, t1.Point.Y, t1.Vector.X, t1.Vector.Y)
}
