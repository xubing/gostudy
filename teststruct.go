package main

import (
	"fmt"
	// "regexp"
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

type CHand struct {
	H string `H`
	K string `k`
}

type SonHand struct {
	Hand
	Age int
}

func (t *Test) ToString() string {
	fmt.Println("Name->", t.Name)
	return "1234"
}
func (t Test) PP() {
	fmt.Println("PP:", t.Name)
}

type KV map[string]string

func main() {
	fmt.Println("test")

	sh := SonHand{Hand{"hello", "son"}, 32}
	fmt.Println(sh)
	//	fmt.Println(sh.H,sh.Age,sh.K)
	//	fmt.Println(sh.Hand)
	//	fmt.Println(&sh,&sh.H,&sh.K,&sh.Hand)
	//	hp := (&sh.H)
	//	fmt.Println((&sh.H),&hp)

	test := Test{"test"}
	fmt.Println(test)
	ret := test.ToString()
	fmt.Println(ret)
	test.PP()

	//	fmt.Println(test.ToString())

	kv := make(KV, 2)
	kv["1"] = "a"
	kv["2"] = "b"
	fmt.Println(kv, kv["1"])

	// h := Hand{"Hello", "Kate"}
	// ret, err := h.(*CHand)
	// if err == nil {
	// 	fmt.Println(ret)
	// }

}
