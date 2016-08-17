// gofunc
package main

import (
	"fmt"
	//	"math"
	//	"time"
)

func add(x, y int) (sum int) {
	return x + y
}

func printRang(max int) {
	for i := 1; i < max; i++ {
		fmt.Println(i)
	}
}

func WordCount(s string) map[string]int {

	sLen := len(s)
	fmt.Println("len:", len(s))
	fmt.Println("ss:", string(s[1]), s[2], s[3])
	var ret = map[string]int{}
	for index := 0; index < sLen; index++ {
		ret[string(s[index])] = ret[string(s[index])] + 1
	}
	return ret
}

func ScanString(str string) {

	for k, v := range str {
		fmt.Println(k, string(v))

	}
}

type BB (int)

type AA struct {
	Age int
}

func undifinedParas(args ...int) {

}

//test return ...
func testRet(max int) (ret int, err error) {
	return max, nil
}

//测试go的function
func main() {

	_, ok := testRet(1)
	fmt.Println(ok)
	var b BB
	b = 10
	aa := struct {
		Age int
	}{1}

	var unnameFuc = func(str string) {
		fmt.Println(str + " index")
	}

	func(str string) {
		fmt.Println(str + " index")
	}("123")
	unnameFuc("Kiss You")
	// var tt = b.(type)
	// fmt.Println(tt)
	fmt.Println(b, aa.Age)
	ScanString("bing")
	var ret = WordCount("123456432432")
	fmt.Println(ret)

	//	fmt.Println(math.Pi * 2)
	//	fmt.Println("Hello China!")
	//	sum := add(1, 3)
	//	fmt.Println(sum)
	//	printRang(10)

	//	today := time.Now().Weekday()
	//	fmt.Println(today)
	//	switch time.Tuesday {
	//	case today + 0:
	//		fmt.Println("Today.")
	//	case today + 1:
	//		fmt.Println("Tomorrow.")
	//	case today + 2:
	//		fmt.Println("In two days.")
	//	default:
	//		fmt.Println("Too far away.")
	//	}

	//	t := time.Now()
	//	switch {
	//	case t.Hour() < 12:
	//		fmt.Println("Good morning!")
	//	case t.Hour() < 17:
	//		fmt.Println("Good afternoon.")
	//	default:
	//		fmt.Println("Good evening.")
	//	}

}
