package main

import (
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Listening on :8080")

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		fmt.Println("Received log from: " + conn.RemoteAddr().String())

		go handleLog(conn)
	}
}

func handleLog(conn net.Conn) {
	defer conn.Close()

	for {

	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	log := string(buf[:reqLen])
	fmt.Println("Received log: " + log)

	fmt.Println("Sending log: server received log message")
	_, err = conn.Write([]byte("server received log message"))
	if err != nil {
		fmt.Println("Error sending response:", err.Error())
	}
}
}
