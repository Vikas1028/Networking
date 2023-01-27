package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// establish connection
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error dialing:", err.Error())
		return
	}
	defer conn.Close()

	// send log every 5 minutes
	for {
		log := "client sen log message"
		fmt.Println("Sending log: " + log)
		conn.Write([]byte(log))

		// receive response
		buf := make([]byte, 1024) 
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		response := string(buf[:reqLen])
		fmt.Println("Received response: " + response)

		time.Sleep(5 * time.Minute)
	}
}
