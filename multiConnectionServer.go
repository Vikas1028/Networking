package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}

	defer ln.Close()

	fmt.Println("Listening on :8000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		wg.Add(1)
		go handleConnection(conn, &wg)
	}
}

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {

	fmt.Println("Connection established")

	for i := 1; i < 6; i++ {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		data := buffer[:n]
		fmt.Println("Received:", string(data))

		_, err = conn.Write([]byte(fmt.Sprintf("hello, %s", data)))
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		time.Sleep(10 * time.Second)
	}
	fmt.Println("Connection closed")
	wg.Done()
	conn.Close()
}
