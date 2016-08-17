package main

import "fmt"

func PT(pt interface{}) {
	fmt.Println(pt)
}

//const{
//size int  =10
//}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

var name = 123

func main() {

	fmt.Println(1 << 0x0000)
	fmt.Println(0x0000 >> 1)
	fmt.Println(1 << 0x0001)
	//	fmt.Println( 1 >> 0x0001 )

	const (
		c0 = iota
		c1 = iota
		c2 = iota
	)

	var ch = make(chan int)
	ch <- 1
	<-ch

	//	fmt.Println(c0)
	//	fmt.Println(c1)
	//	fmt.Println(c2)
	//	fmt.Println( 1<< c0)
	//	fmt.Println( 1<< 0)
	//	fmt.Println( 1<< 1)
	const (
		d0 = 1 << iota
		d1 = 1 << iota
		d2 = 1 << iota
	)

	//	fmt.Println(d0)
	//	fmt.Println(d1)
	//	fmt.Println(d2)

}
