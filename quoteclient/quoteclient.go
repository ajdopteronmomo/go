package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.10.30.155:9129")

	if err != nil {
		fmt.Println("dial failed", err)
		os.Exit(1)
	}

	defer conn.Close()

	buffer := make([]byte, 512)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("read failed", err)
		return
	}

	fmt.Println("count:", n, "msg:", string(buffer))
}
