package main

import (
	"fmt"
	"net"
)

// 客户端

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	conn.Write([]byte("hello"))
	conn.Close()
}
