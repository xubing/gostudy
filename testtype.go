package main

import (
	"fmt"
	//	"math"
)

const (
	a = 1
	b
	c
)

const (
	d, e = 1, 2
	f, g
)

const (
	aa = 2
	bb = iota
	cc = iota
	dd = 0
	ee = iota
)

const (
	cccc = 100
	bbbb = iota
)

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
)

func main() {

	p := new([10]int)
	p[2] = 2
	fmt.Println(p)
}
