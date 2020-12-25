package main

// go fmt 格式化
import (
	"fmt"
)

func _for() {
	// 当i=5时就跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 跳出for循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// 当i=5时，跳过此次for循环（不执行for循环内部的打印语句），继续下一次循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // 继续下一次循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}

func _goto() {
	// 跳出多层for循环
	var flag = false
	for i := 0; i < 10; i++ {
		// 单引号: 字符
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break // 跳出内层的for循环
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break // 跳出for循环（外层的for循环）
		}
	}

	// goto+label实现跳出多层for循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto A // 跳到我指定的那个标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
A: // label标签
	fmt.Println("over")
}

func _switch() {
	a := 2
	switch {
	case a == 1:
		fmt.Println("1")
	default:
		fmt.Println("0")
	}

	// 多个判断
	switch a {
	case 1, 5, 3, 7:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	}
	// 内部声明
	switch n := 3; n {
	case 1, 5, 3, 7:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	}
}

func _op() {
	var (
		a uint8 = 2
		b uint8 = 5
	)
	// 关系运算符
	fmt.Println(a == b) // Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b) // 不等于

	// 位运算:针对的是二进制数
	// 5的二进制表示：101
	// 2的二进制表示： 10

	// &:按位与 (两位均为1才为1)
	fmt.Println(5 & 2) // 000
	// |:按位或（两位有1个为1就为1）
	fmt.Println(5 | 2) // 111
	// ^:按位异或（两位不一样则为1）
	fmt.Println(5 ^ 2) // 111
	// <<:将二进制位左移指定位数
	fmt.Println(5 << 1)  // 1010 => 10
	fmt.Println(1 << 10) // 10000000000 => 1024
	// >>:将二进制位右移指定的位数
	fmt.Println(5 >> 2)  // 1 =>1
	var m = int8(1)      // 只能存8位 -> 127
	fmt.Println(m << 10) // 10000000000

	var x int
	x = 10
	x++    // x = x + 1
	x--    // x = x - 1
	x *= 2 // x = x * 2
	x /= 2 // x = x / 2
	x %= 2 // x = x % 2

	x <<= 2 // x = x << 2
	x &= 2  // x = x & 2
	x |= 3  // x = x | 3   ->  1 1 -> 3
	x ^= 4  // x = x ^ 4	 ->1 0 0 -> 1 1 1 ->  7
	x >>= 2 // x = x >> 2  -> 1

	fmt.Println(x)
}

// 入口函数
func main() {
	fmt.Println("break continue:")
	_for()
	fmt.Println()
	fmt.Println()
	fmt.Println("goto:")
	_goto()
	fmt.Println()
	fmt.Println("switch:")
	_switch()
	fmt.Println()
	fmt.Println("运算符:")
	_op()
}
