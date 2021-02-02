package main

import (
	"fmt"
	"image"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var x int64 = 0         // 公共资源
var lock sync.Mutex     // 锁住公共资源
var rwlock sync.RWMutex // 读写锁
var once sync.Once      // once.Do(func())确保只执行一次 ==> eg: 并发, 初始化, 关闭channel !! 需要传递参数就需要搭配闭包来使用。

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

// 读写互斥锁 读写锁非常适合读多写少的场景
// RWMutex是读写互斥锁。该锁可以被同时多个读取者持有或唯一个写入者持有。RWMutex可以创建为其他结构体的字段；零值为解锁状态。RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁。
func write() {
	// lock.Lock() // 加互斥锁
	rwlock.Lock()
	x++
	time.Sleep(10 * time.Millisecond) // 假设写操作耗时10毫秒
	// lock.Unlock()
	rwlock.Unlock()
	wg.Done()
}
func read() {
	// lock.Lock()
	rwlock.RLock()
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	// lock.Unlock()
	rwlock.RUnlock()
	wg.Done()
}

func rw() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

// sync.Once
var icons map[string]image.Image

var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":  Icon("left.png"),
		"up":    Icon("up.png"),
		"right": Icon("right.png"),
		"down":  Icon("down.png"),
	}
}

// Icon 是并发安全的
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	// 这样被多个goroutine调用时不是并发安全的
	// if icons == nil {
	// 	loadIcons()
	// }
	return icons[name]
}

// sync.Map  并发访问map  赋值: Store(key, v) 取值Load(key)  LoadOrStore、删除: Delete、遍历: Range
var m = sync.Map{}

func _syncMap() {
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("遍历:")
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

// atomic  原子操作并发安全且比互斥锁效率更高
func _atomic() {
	var a int64 = 0
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&a, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(atomic.LoadInt64(&a), end.Sub(start))
	// CompareAndSwapInt64(&p, old, new) (bool) -> old == 原值的话与新值交换, 返回是否交换
}

func main() {
	fmt.Println()
	fmt.Println("sync:")
	_sync()
	fmt.Println()
	fmt.Println("读写锁测试耗时:")
	rw()
	/*
		读写: ~107.0104ms
		互斥锁: ~1.5765028s
	*/
	fmt.Println()
	fmt.Println("sync.Once:")
	rw()
	// Go语言中内置的map不是并发安全的
	fmt.Println()
	fmt.Println("sync.Map:")
	_syncMap()
	// 原子操作 atomic
	fmt.Println()
	fmt.Println("sync/atomic:")
	_atomic()

}
