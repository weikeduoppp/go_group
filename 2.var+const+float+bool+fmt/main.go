package main

// go fmt 格式化
import (
	"fmt"
	"math"
)

//  Go语言中标识符由字母数字和_(下划线）组成，并且只能以字母和_开头。 举几个例子：abc, _, _123, a123。

// 函数外只能放置(变量/常量/函数/类型)的声明
var a string

// 批量声明  需要事先声明 驼峰命名方式
var (
	number int
	isNum  bool
)

func foo() (int, bool) {
	return 9, true
}

// 非全局变量声明了需要使用
func _var() {
	number = 1
	isNum = true
	// noGobal := 2  // no use
	var name string = "san"
	message := "213"            // :=只能函数内部使用  -- 短变量声明
	fmt.Print(number)           // 打印
	fmt.Printf("name:%v", name) // %v占位符.
	fmt.Println(message)        // Println打印完加换行符

	// _ -> 匿名变量. 需要忽略某变量
	x, _ := foo()
	fmt.Println(x)
}

// 常量
const pi = 3.14

// 多个常量
const (
	SUCCESS = 200
	FAIL    = 400
)

// 省略了值则表示和上面一行的值相同 n1 = n2 = n3 = 100
const (
	p1 = 100
	p2
	p3
	p4
)

// iota: 常量计数器
// iota在const关键字出现时将被重置为0。
// ** const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。**
// 使用iota能简化定义，在定义枚举时很有用。

const (
	n1 = iota //0
	n2        //1
	n3        //2
	n4        //3
)

// 几个常见的iota示例:
// 使用_跳过某些值
const (
	a1 = iota // 0
	_         // 1
	a2 = iota // 2
	a3 = iota // 3
)

// iota声明中间插队
const (
	b1 = iota // 0
	b2 = 100
	b3 = iota // 2
	b4        // 3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // 1024 (2^10)
	MB = 1 << (10 * iota) // 1024 * 1024 ( 2^20 )
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func _const() {
	fmt.Println("n1: ", n1)
	fmt.Println("n2: ", n2)
	fmt.Println("n3: ", n3)
	fmt.Println("n4: ", n4)
	fmt.Println("b1: ", b1)
	fmt.Println("b2: ", b2)
	fmt.Println("b3: ", b3)
	fmt.Println("b4: ", b4)
}

// 浮点数 float32 float64
func _floatAndComplex() {
	fmt.Println(math.MaxFloat32)
	f1 := 1.223
	fmt.Printf("%T\n", f1) // Go默认float64.
	f2 := float32(1.214)
	fmt.Printf("%T\n", f2) // 显示声明

	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
}

func _bool() {
	// bool默认false.  数值不能强制转换布尔型.  不参与数值运算
	var b1 bool // false
	b2 := true
	fmt.Printf("%T\n", b2) // bool
	fmt.Println(b1)
	fmt.Println(b2)
}

// fmt占位符
func _fmt() {
	n := 100
	fmt.Printf("%T\n", n) // 查看类型
	fmt.Printf("%v\n", n) // 查看值
	fmt.Printf("%b\n", n) // 查看二进制
	fmt.Printf("%d\n", n) // 查看十进制
	fmt.Printf("%o\n", n) // 查看八进制
	fmt.Printf("%x\n", n) // 查看十六进制

	s := "hello"
	fmt.Printf("%s\n", s)  // 查看字符串
	fmt.Printf("%#v\n", s) // 加# 多双引号 查看值
}

// 入口函数
func main() {
	fmt.Println("变量:")
	_var()
	fmt.Println()
	fmt.Println("常量:")
	_const()
	fmt.Println()
	fmt.Println("浮点数:")
	_floatAndComplex()
	fmt.Println()
	fmt.Println("布尔值:")
	_bool()
	fmt.Println()
	fmt.Println("fmt:")
	_fmt()
}
