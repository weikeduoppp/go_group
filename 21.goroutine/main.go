package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// sync.WaitGroup
var wg sync.WaitGroup

func _func(i int) {
	fmt.Println("func", i)
}
func _func2(i int) {
	defer wg.Done()
	fmt.Println("func", i)
}
func _goroutine() {
	fmt.Println()
	fmt.Println("goroutine:")
	for i := 0; i < 2; i++ {
		go _func(i) // 单独开启一个goroutine  创建新的goroutine的时候需要花费一些时间
	}
	time.Sleep(time.Second)
}

func _wait() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go _func2(i) // 单独开启一个goroutine  创建新的goroutine的时候需要花费一些时间
	}
	wg.Wait()
}

// math/rand 随机数
func _rand() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10) /// [0,n)的随机数
		fmt.Println(r1, r2)
	}
}

func main() {
	_goroutine()
	// main 就是一个goroutine.  main结束了. 其中创建的goroutine也会结束
	_rand()
	_wait()
}
