package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var x int64 = 0     // 公共资源
var lock sync.Mutex // 锁住公共资源

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 写的锁
		x++
		lock.Unlock()
	}
	wg.Done()
}

func _sync() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

// 读写互斥锁
// RWMutex是读写互斥锁。该锁可以被同时多个读取者持有或唯一个写入者持有。RWMutex可以创建为其他结构体的字段；零值为解锁状态。RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁。

func main() {
	fmt.Println()
	fmt.Println("sync:")
	_sync()

}
