// testchannel
package main

import "fmt"

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser = MyFloat(-3.453)

	// f := MyFloat(-3.453)
	// a = f

	fmt.Println(a)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
