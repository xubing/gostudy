package main

import (
	"fmt"
	"log"
	"net/http"
)

const SERVER_PORT = 8081
const SERVER_DOMAIN = "localhost"
const RESPONSE_TEMPLATE = "hello"

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {

	//	http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
	//	http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "ca.crt", "server.key", nil)

	//	h := http.FileServer(http.Dir("."))
	//	http.ListenAndServeTLS(":8001", "rui.crt", "rui.key", h)

	//	fmt.Println(SERVER_DOMAIN,SERVER_PORT,RESPONSE_TEMPLATE)
	//	h := http.FileServer(http.Dir("."))
	//	http.ListenAndServeTLS(":8001","ca.crt","server.key",h)

	http.HandleFunc("/", SayHello)
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
