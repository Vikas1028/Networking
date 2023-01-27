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

func handleConnection(conn net.Conn) {

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	handleError(err)

	fmt.Println("Received message: ", string(buf[:n]))

	_, err = conn.Write([]byte("Hello, Client"))

}

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	handleError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		handleError(err)
		go handleConnection(conn)
	}
}
