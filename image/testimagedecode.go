package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	//	"bufio"
)

func main() {
	flag.Parse()
	//	rand.Seed(time.Now().UTC().UnixNano())

	filename := "/Users/Bing/Desktop/picdir3/benchRGB.png"
	reader, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("decoding image")
	_, _, err = image.Decode(reader)
	if err != nil {
		fmt.Println("image.Decode fail: ", err.Error())
	}
	defer reader.Close()

}
