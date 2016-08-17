package main

type Base struct {
	Name string
}

type Foo struct {
	Base
}

func main() {

	b := Base{"xxx"}
	f := Foo{Base{"foo"}, 30}

}
