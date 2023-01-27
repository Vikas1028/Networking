package main

import (
	"fmt"
	"log"
	"net"
)

func handleError(err error) {
	if err != nil {
		log.Fatalf("Err: ", err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	handleError(err)
	defer conn.Close()

	for {

		var data string
		fmt.Print("Enter text: ")
		fmt.Scanln(&data)
		fmt.Fprintf(conn, data+"\n")
	}
}
