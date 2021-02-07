package main

import (
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	var temp = [128]byte{}
	n, err := conn.Read(temp[:])
	if err != nil {
		fmt.Printf("conn read err: %v", err)
	}
	fmt.Println(string(temp[:n]))
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("tcp err: %v", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		fmt.Printf("listen err: %v", err)
	}
	handleConn(conn)
}
