package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	//	"net"
	"io/ioutil"
	"strings"
)

var version = "v1.001.201605171500"
var size = 0
var port = ":8002"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloHandler...")
	io.WriteString(w, "12000000")
}
func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("versionHandler...")
	io.WriteString(w, version)
}

func main() {
	//	addrs, err := net.InterfaceAddrs()
	//	fmt.Println("address:",addrs)

	files, err := ioutil.ReadDir("./server/")
	if err != nil {
		log.Fatal(err)
	}

	var fileZip = "v0.0.0"
	for _, file := range files {
		filename := file.Name()
		if filename[0] == 'v' {
			//			fmt.Println(filename)
			if strings.Compare(filename, fileZip) > 0 {
				fileZip = filename
			}
		}
	}

	if fileZip == "" {
		log.Fatal("不存在更新包")
	}
	version = strings.Replace(fileZip, ".zip", "", 1)

	fmt.Println("\033[6;32;40m available package:" + version + "\033[0m")
	fmt.Println("\033[6;3;40m start server at: 127.0.0.0" + port + "\033[0m")
	http.HandleFunc("/size", helloHandler)
	http.HandleFunc("/version", versionHandler)
	http.Handle("/", http.FileServer(http.Dir("./server")))
	err = http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}
