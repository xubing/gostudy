package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	//request a url.
	resp, err := http.GET("http://www.baidu.com")
	if err != nil {
		fmt.Println(resp)
		ddreturn
	}

	//	io.Copy(os.Stdout,resp.Body)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

}
