package main

// go fmt 格式化
import (
	"fmt"
)


func _if() {
	// 作用域 age 单独声明再 if{}
	if age := 19; age > 18 {
		fmt.Println("大于18")
	} else {
		fmt.Println("否则")
	}
	// fmt.Println(age) // err
}
func _for() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 2
	a := 5
	for ; a < 10; a++ {
		fmt.Println(a)
	}

	// 3
	b := 5
	for b < 10 {
		fmt.Println(b)
		b++
	}
}
func _forrange() {
	str := "hello傻字123"
	for i,r := range str {
		// r: ASCII码
		fmt.Printf("%v %c\n", i, r)
	}
}

func _ninenine() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%v✖️%v=%v ",i,j,i*j)
		}
		fmt.Println()
	}
}


// 入口函数
func main() {
	fmt.Println("if流程控制:")
	_if()
	fmt.Println()
	fmt.Println()
	fmt.Println("for循环:")
	_for()
	fmt.Println()
	fmt.Println()
	fmt.Println("forrange:")
	_forrange()
	fmt.Println()
	fmt.Println("九九法:")
	_ninenine()
	// 中文字符 一般长度为3字节
	fmt.Println(len("九九"))
}
