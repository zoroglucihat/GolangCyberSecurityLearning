package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "go.com:80")
	if err == nil {
		fmt.Println("Connection successful")
	}
}
