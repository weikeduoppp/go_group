package main

import (
	"fmt"
	"log"
)

func _func() {
	// 设置输出位置 默认终端
	// log.SetOutput(os.Stdout)
	// log 简单的日志服务
	log.Println("输出")
}

func main() {
	fmt.Println()
	fmt.Println("log:")
	_func()
}
