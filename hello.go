package main

import (
	"fmt"
	"regexp"
)

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

type Test struct {
	Name string
}

type Hand struct {
	H string `H`
	K string `k`
}

func (t *Test) ToString() {
	fmt.Println("Name->", t.Name)
}

//TODO 如何修改参数，支持任意长度的数组。
func modifyArray(pArray *[10]int) {
	pArray[1] = 10
}

func modifySlice(slice []int) {
	slice[1] = 20
}

type KV map[string]string

func main() {
	str := "hello中国"
	fmt.Println("len:", len(str))
	for i, ch := range str {
		fmt.Println(i, string(ch))
	}

	type Animal struct {
		Name  string
		Order string
	}
	var animals = []Animal{{"Platypus", "Monotremata"}}
	fmt.Println(animals)

	var name int = 8
	var name1 rune
	name1 = 23
	fmt.Println("name", name, name1)
	fmt.Println("test")
	var tt = make([]int, 1)
	fmt.Println(tt)
	var kv = make(KV)
	kv["1"] = "1*1"
	fmt.Println(kv)
	delete(kv, "111")
	fmt.Println(kv)

	//
	pArray := &[10]int{1, 3, 5}
	var slice = (*pArray)[:]
	modifySlice(slice)
	fmt.Println(slice)
	fmt.Println(pArray)

	modifyArray(pArray)
	fmt.Println(pArray)

	ArrayInt := [5]int{}
	for k, v := range ArrayInt {
		fmt.Println(k, v)
	}

	fmt.Println(80 % 4)
	var h = Hand{"HH", "KK"}
	fmt.Println("hh->", h.H, h.K)
	fmt.Println("Hello world Bing")

	matched, err := regexp.MatchString("foo.*", "seafood")
	fmt.Println(matched, err)

	matched, err = regexp.MatchString("a(b", "seafood")
	fmt.Println(matched, err)

	var array = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println()

	var t = Test{"Bing"}
	t.ToString()
	// for i := 0; i < len(array); i++ {
	// 	fmt.Println(array[i])
	// }

	for k, v := range array {
		fmt.Println(k, v)
	}

	var a []int
	printSlice("a", a)

}
