package main

// go fmt 格式化
import (
	"fmt"
)

/*
	new 分配内存, 主要分配值类型, 返回指针类型
	使用 panic/recover 模式处理错误
*/
func _func() {
	// panic: 程序崩溃
	f1()
	f2()
}

func f1() {
	defer func() {
		/*
			recover()必须搭配defer使用。
			defer一定要在可能引发panic的语句之前定义。
		*/
		err := recover()
		fmt.Println(err)
	}()
	panic("程序崩溃")
	fmt.Println("1")
}

func f2() {
	fmt.Println("2")
}

func _fmt() {
	/*
		Printf
		%T 类型
		%d 十进制
		%b 二进制
		%o 八进制
		%x 十六进制
		%c 字符
		%p 指针
		%f 浮点数
		%t 布尔值
		%% 百分号
	*/

	// 获取输入 fmt.Scan() Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符。
	var (
		name    string
		age     int
		married bool
	)
	// fmt.Scan(&name, &age, &married)
	// fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

// 入口函数
func main() {
	fmt.Println()
	fmt.Println("内置函数:")
	_func()
	fmt.Println()
	fmt.Println("fmt:")
	_fmt()
}
