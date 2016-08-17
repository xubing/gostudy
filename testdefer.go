package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}
func un(s string) {
	fmt.Println("leaving:", s)
}
func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}
func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func CM(m map[string]string) {
	m["a"] = "a"
}

func main() {
	//	b()

	//	a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	//	fmt.Println(a)

	m := map[string]string{"a": "aa", "b": "bb"}
	fmt.Println(m)
	CM(m)

	m := map[string]string{"a": "aa", "b": "bb"}
	a := m["a"]
	b, err = m["b"]
	fmt.Println(a)

	//	var b string
	//	var err bool
	b, err := m["a"]
	c := m["c"]

	if err {
		fmt.Println(b, c)
	}

}
