package main

import (
	"fmt"
	//	"hash"
	"crypto/md5"
)

func main() {

	fmt.Println("Test hash")
	teststr := "testhash1"

	MD5Instance := md5.New()
	MD5Instance.Write([]byte(teststr))
	Result := MD5Instance.Sum([]byte(""))
	fmt.Printf("%x\n", Result)
	ar := fmt.Sprintf("%x", Result)
	fmt.Println(ar)

}
