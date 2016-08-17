package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func checkError(err error) {

	if err != nil {
		fmt.Println("error:  " + err.Error())
		os.Exit(1)
	}
}

func checkSum(msg []byte) uint16 {
	sum := 0
	for n := 1; n < len(msg); n += 2 {

		sum += int(msg[n])*256 + int(msg[n+1])
	}

	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {

		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func main() {

	fmt.Println("Test Simple Http. Head")

	if len(os.Args) != 2 {
		fmt.Println("Uage: ", os.Args[0], "host")
		os.Exit(1)
	}

	service := os.Args[1] //"www.baidu.com"
	conn, err := net.Dial("tcp", service)
	checkError(err)

	_, err = conn.Write([]byte("HEAD/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println("Got response")
	fmt.Println(result)
	fmt.Println(string(result))

	os.Exit(0)

}
