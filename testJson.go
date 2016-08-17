package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type DD struct {
	Age  int
	Name string
}
type Book struct {
	Title       string
	Author      []string
	Publisher   string
	IsPublisher bool
	Price       float32
	More        *DD
}

func testDecode() {

	const jsonStream = `
		[
			{"Name": "Ed", "Text": "Knock knock."},
			{"Name": "Sam", "Text": "Who's there?"},
			{"Name": "Ed", "Text": "Go fmt."},
			{"Name": "Sam", "Text": "Go fmt who?"},
			{"Name": "Ed", "Text": "Go fmt yourself!"}
		]
	`
	type Message struct {
		Name, Text string
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("T: %T: %v\n", t, t)

	var m Message
	// while the array contains values
	for dec.More() {

		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.Name, m.Text)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

}

func main() {

	testDecode()

	d := DD{20, "DD"}
	gobook := Book{"Go lang",
		[]string{"111", "222"},
		"Publ1",
		true,
		32.2, &d}
	b, err := json.Marshal(gobook)

	fmt.Println(gobook)
	if err != nil {
		fmt.Println("Parse error")
	}
	fmt.Println(string(b))

	//解析

	book := []byte(`{"Title":"Go lang","Author":["333","222"],"Publisher":"Publ1","IsPublisher":true,"Price":32.2,"More":{"Age":20,"Name":"DD"}}`)
	var bookun Book
	err = json.Unmarshal(book, &bookun)
	fmt.Println(string(book))

	//解析未知结构的Json
	book2 := []byte(`{"Title":"Go lang","Author":["333","222"],"Publisher":"Publ1"}`)
	var b2 interface{}
	if err = json.Unmarshal(book2, &b2); err == nil {
		fmt.Println("unkonwn:", b2)
	} else {
		fmt.Println("xxxx", err.Error())
	}

}
