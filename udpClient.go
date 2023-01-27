package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	i:=0
	for {
        
		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
		i++
		conn.Write([]byte("data packate: " + string(i)))
		time.Sleep(10 * time.Second)
	}
}
