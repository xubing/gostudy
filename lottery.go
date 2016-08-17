package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Start")
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":8080", nil)
}
