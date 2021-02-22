package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"25.protocol/proto"
)

// 粘包
func handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("conn read err: %v\n", err)
			return
		}
		fmt.Println("收到:", string(msg))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("tcp err: %v", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("listen err: %v", err)
			return
		}
		go handleConn(conn)
	}
}
