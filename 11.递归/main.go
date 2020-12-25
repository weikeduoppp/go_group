package main

// go fmt 格式化
import (
	"fmt"
)

func _func() {
	fmt.Println(f1(5))
}

// 阶乘
func f1(n uint32) uint32 {
	if n == 1 {
		return 1
	}
	return n * f1(n-1)
}

// 面试题.
/*
	假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
	每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/
func f2(n int) int {
	if n <= 2 {
		return n
	}
	var (
		i int = 1
		j int = 2
	)
	for k := 3; k <= n; k++ {
		i, j = j, i+j
	}
	return j
}

// 入口函数
func main() {
	fmt.Println()
	fmt.Println("阶乘:")
	_func()
	fmt.Println()
	fmt.Println("上台阶:")
	fmt.Println(f2(4))
}
