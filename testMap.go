package main

import (
	"fmt"
	//	 "strconv"
	"unsafe"
)

func main() {

	fmt.Println("Test map")
	var m = make(map[int]int, 5)
	for index := 0; index < 10; index++ {
		m[index] = index
		fmt.Println(m[index])
	}
	fmt.Println(m, len(m))
	var v, ok = m[20]
	fmt.Println(ok, v)

	m[20]++
	v, ok = m[20]
	fmt.Println(ok, v)

	//	var i int
	//	fmt.Scanf("%d", &i)
	//	str := strconv.FormatInt(int64(i), 10)
	//	hex, _ := strconv.ParseInt(str, 16, 64)
	//	fmt.Printf("%d\n", hex)

	str := "A Go variable"
	addr := unsafe.Pointer(&str)
	fmt.Printf("The address of str is %d\n", addr)
	str2 := (*string)(addr)
	fmt.Printf("String constructed from pointer: %s\n", *str2)
	address := uintptr(addr)
	address += 4
	// This has undefined behavior!
	str3 := (*string)(unsafe.Pointer(address))
	fmt.Printf("String constructed from pointer:%s\n", *str3)

}
