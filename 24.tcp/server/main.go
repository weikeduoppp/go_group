package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		var temp = [128]byte{}
		n, err := conn.Read(temp[:])
		if err != nil {
			fmt.Printf("conn read err: %v\n", err)
			return
		}
		fmt.Println(string(temp[:n]))
		// 回复客户端
		fmt.Println("请回复:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" {
			break
		}
		conn.Write([]byte(text))
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
