package main

import (
	"fmt"
	"net"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {

		buffer := make([]byte, 4096)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleData(buffer[:n])
	}
}

func handleData(data []byte) {

	fmt.Println(string(data))
}
