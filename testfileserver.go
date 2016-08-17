package main

import (
	"fmt"
	// "io"
	"log"
	"net/http"
	//	"net"
	// "io/ioutil"
	// "strings"
)

var port = ":8002"

func main() {

	fmt.Println("\033[6;3;40m start server at: localhost:" + port + "\033[0m")
	http.Handle("/", http.FileServer(http.Dir("./server")))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}
