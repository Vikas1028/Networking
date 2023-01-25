package main

import (
	"fmt"
	"log"
	"net"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	handleError(err)
	defer conn.Close()

	_, err = conn.Write([]byte("Hello, remote Server"))
	handleError(err)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	handleError(err)

	fmt.Println("Received message: ", string(buf[:n]))
}
