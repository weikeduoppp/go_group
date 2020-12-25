package main

import "fmt"

// 函数外只能放置(变量/常量/函数/类型)的声明
var a string
// 非法
// fmt.Println(a)

// 入口函数
func main() {
	message := "213"
	fmt.Println(message)
}
