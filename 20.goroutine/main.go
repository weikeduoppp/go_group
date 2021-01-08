package main

import (
	"fmt"
	"time"
)

func _func(i int) {
	fmt.Println("func", i)
}

func main() {
	fmt.Println()
	fmt.Println("goroutine:")
	for i := 0; i < 100; i++ {
		go _func(i) // 单独开启一个goroutine
	}
	time.Sleep(time.Second)
	// main 就是一个goroutine.  main结束了. 其中创建的goroutine也会结束
}
