package main

import (
	"fmt"
)

/*
#include <stdlib.h>
*/
import "C"

func Random() int {
	return int(C.random())

}

func main() {

	ret := Random()

	fmt.Println(ret)

}
