package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error dialing:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	message := []byte("client 1")

	for i := 1; i < 6; i++ {
		_, err = conn.Write(message)
		if err != nil {
			fmt.Println("Error sending data:", err.Error())
			return
		}

		buffer := make([]byte, 4096)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error receiving data:", err.Error())
			return
		}
		fmt.Println("Data received:", string(buffer[:n]))
	}

	fmt.Println("Disconnected from server")
}
