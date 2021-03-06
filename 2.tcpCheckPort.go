package main

import (
	"fmt"
	"net"
)

func main() {

	for i := 1; i <= 65535; i++ {
		address := fmt.Sprintf("192.168.1.105:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
			//port kapalı
		}
		conn.Close()
		fmt.Printf("%d open \n", i)

	}
}
