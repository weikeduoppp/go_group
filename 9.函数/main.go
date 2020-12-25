package main

// go fmt 格式化
import (
	"fmt"
	"strings"
)

func sum(x int, y int) (ret int) {
	ret = x + y
	return // 有命名return后可省略ret
}

func multiple() (x int, y int) { // 可多个返回值
	return 1, 3
}

// 多个相同类型可以省略. 只写最后
func multipleParams(x, y, z int) { // 可多个返回值
}

// 可选可变长
func optional(y ...int) {
	fmt.Println(y) // 切片slice  type: []int
}

func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	}
	return nil
}

// 传参都是拷贝(副本).
func _func() {
	fmt.Println(sum(1, 2))
	x, y := multiple()
	fmt.Println(x, y)
	optional()
	optional(1, 2, 3)
}

// 作用域: 全局. 函数内部. 语句块(if, for).

var num int = 1

// defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行
// defer执行时机: 在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而defer语句执行的时机就在返回值赋值操作后，RET指令执行前
// defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
func _defer() {
	defer func() {
		num++
	}()
	defer fmt.Println(111)
	defer fmt.Println(222) // 多个defer 先进后出

	// 用于 资源清理、文件关闭、解锁及记录时间 , websocket

	// defer 经典案例
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x // ret = 5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // ret = x = 5  x++  return 6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // ret = y = x = 5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // x = 5 函数传参是副本 不影响ret
}
func f5() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5 // x = 5 函数传参是x内存地址 修改是x内存地址的值, ret=x=6
}

// 声明函数类型
type calculation func() int

func _funcType() {
	var c calculation
	c = f1
	fmt.Printf("%T\n", c)
	b := f1
	fmt.Printf("%T\n", b)
}

// 1.参数是函数 2.定义匿名函数
func f6(x func() int) func() int {
	// 函数内部不能声明带名字的函数
	var f1 = func(i int) int { return i } // 匿名函数
	f1(2)
	return x
}

//  闭包指的是一个函数和与其相关的引用环境组合而成的实体
func adder(x int) func(int) int {
	// 在生命周期内 变量x也一直有效
	return func(y int) int {
		x += y
		fmt.Println(x)
		return x
	}
}

// test func1(func2) // eg: 实现把需要传递两个参数的函数包装成不需要传递参数的函数
func func1(f func()) {
	fmt.Println("func1")
	f()
}
func func2(x, y int) {
	fmt.Println(x, y)
}

func func3(f func(int, int), x, y int) func() {
	return func() {
		f(x, y)
	}
}

/*
	你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
	分配规则如下：
	a. 名字中每包含1个'e'或'E'分1枚金币
	b. 名字中每包含1个'i'或'I'分2枚金币
	c. 名字中每包含1个'o'或'O'分3枚金币
	d: 名字中每包含1个'u'或'U'分4枚金币
	写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
	程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
func _test() {
	var (
		coins = 50
		users = []string{
			"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
		}
		distribution = make(map[string]int, len(users))
	)
	/*
		1.先切割每个名字. 掌握每个字母包含几个
		2. 依次判断分配金币. 统计分配出去的金币
		3. 得出剩余
	*/
	for _, v := range users {
		slice := strings.Split(v, "")
		_map := findLetter(slice)
		distribution[v] = 0
		if _map["e"] != 0 {
			distribution[v] += _map["e"] * 1
			coins -= _map["e"] * 1
		}
		if _map["E"] != 0 {
			distribution[v] += _map["E"] * 1
			coins -= _map["E"] * 1
		}
		if _map["i"] != 0 {
			distribution[v] += _map["i"] * 2
			coins -= _map["i"] * 2
		}
		if _map["I"] != 0 {
			distribution[v] += _map["I"] * 2
			coins -= _map["I"] * 2
		}
		if _map["o"] != 0 {
			distribution[v] += _map["o"] * 3
			coins -= _map["o"] * 3
		}
		if _map["O"] != 0 {
			distribution[v] += _map["O"] * 3
			coins -= _map["O"] * 3
		}
		if _map["u"] != 0 {
			distribution[v] += _map["u"] * 4
			coins -= _map["u"] * 4
		}
		if _map["U"] != 0 {
			distribution[v] += _map["U"] * 4
			coins -= _map["U"] * 4
		}
	}
	fmt.Println(coins, distribution)
}
func findLetter(slice []string) map[string]int {
	// 优化: 在这处理分配
	_map := make(map[string]int, len(slice))
	for _, i := range slice {
		if _, ok := _map[i]; !ok {
			_map[i] = 1
		} else {
			_map[i]++
		}
	}
	return _map
}

// 入口函数
func main() {
	fmt.Println()
	fmt.Println("func:")
	_func()
	fmt.Println()
	fmt.Println("defer:")
	_defer()
	fmt.Println(num)
	fmt.Println()
	fmt.Println("_funcType:")
	_funcType()
	fmt.Println()
	fmt.Println("闭包:")
	f := adder(0)
	f(10)
	f(20)
	func1(func3(func2, 2, 3))
	fmt.Println()
	fmt.Println("测试:")
	_test()
}
