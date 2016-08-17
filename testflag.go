package main

import (
	"flag"
	"log"
)

func main() {

	flag.Parse()
	filename := flag.Arg(0)

	log.Println(filename)

}
