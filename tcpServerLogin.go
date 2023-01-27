package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8000")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Response from server:", resp.Status)
}
