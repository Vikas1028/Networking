package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func handleError(err error) {
	if err != nil {
		log.Fatalf("Err: ", err)
	}
}

func handleConnection(conn net.Conn) {

	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		handleError(err)
		log.Printf("[%v] [%v] %v", time.Now(), conn.RemoteAddr(), message)
	}
}

func main() {

	listener, err := net.Listen("tcp", ":8080")
	handleError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		handleError(err)
		go handleConnection(conn)
	}
}
