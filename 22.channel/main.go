package main

import (
	"fmt"
	"time"
)

func _func(i chan int) {
	fmt.Println(<-i)
}

func _func2(i int) {
}

// var name chan T 需要初始化才能使用
func _chan() {
	var a chan int
	a = make(chan int)      // 无缓冲的通道  !! 无缓冲的通道必须有接收才能发送。无缓冲通道也被称为同步通道
	b := make(chan int, 16) // 缓冲的通道
	fmt.Println(a, b)
	/*
		发送: a <- 1
	*/
	go _func(a)
	b <- 1
	a <- 10
	fmt.Println("发送成功a", <-b)
	// close(ch) 通过调用内置的close函数来关闭通道。
	/*
		关闭后的通道有以下特点：

		对一个关闭的通道再发送值就会导致panic。
		对一个关闭的通道进行接收会一直获取值直到通道为空。
		对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
		关闭一个已经关闭的通道会导致panic。
	*/
}

// for range从通道循环取值
func chan1() {
	// i, ok = close(ch) 如果通道关闭 ok为false
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

// 单向通道.  限制操作
func f1(a chan<- int) {} // 只能发送值
func f2(a <-chan int) {} // 只能取值

// WorkerPool 使用可以指定启动的goroutine数量–worker pool模式，控制goroutine的数量，防止goroutine泄漏和暴涨。
func WorkerPool() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

// SelectChannel ...
/*
	- 可处理一个或多个channel的发送/接收操作。
	- 如果多个case同时满足，select会随机选择一个。
	- 对于没有case的select{}会一直等待，可用于阻塞main函数。
*/
func SelectChannel() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: // 有值
			fmt.Println(x)
		case ch <- i: // 没值 ==> 赋值
		}
	}
}

func main() {
	fmt.Println()
	fmt.Println("chan:")
	_chan()
	fmt.Println()
	fmt.Println("range从通道循环取值:")
	chan1()
	fmt.Println()
	fmt.Println("WorkerPool:")
	WorkerPool()
	fmt.Println()
	fmt.Println("select多路复用:")
	SelectChannel()

}
