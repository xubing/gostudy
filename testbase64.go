package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	cnstr := "中国人"
	datecn := []byte(cnstr)
	inputstr := base64.StdEncoding.EncodeToString(datecn)
	fmt.Println(inputstr)

	str := "YWFiYmNjCuaIkeaYr+S4reWbveS6ugoj5LiN5piv"
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%q\n", data)
}
