package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"25.protocol/proto"
)

var wg sync.WaitGroup

func send(conn net.Conn) {
	defer conn.Close()
	// 有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入内容:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "q" {
			wg.Done()
			break
		}
		conn.Write([]byte(text))
	}
}

// 连续发送多条
func sendMore(conn net.Conn) {
	defer conn.Close()
	for i := 0; i < 10; i++ {
		msg := "hello,Hello. How are you?"
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
	defer wg.Done()
}

// 读
func read(conn net.Conn) {
	defer conn.Close()
	for {
		var temp = [128]byte{}
		n, err := conn.Read(temp[:])
		if err != nil {
			fmt.Printf("conn read err: %v", err)
			return
		}
		fmt.Println("服务端回复: " + string(temp[:n]))
	}
}

// 客户端
func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	wg.Add(1)
	go sendMore(conn)
	wg.Wait()

}
