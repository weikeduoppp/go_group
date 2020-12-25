package main

// go fmt 格式化
import (
	"fmt"
	"strings"
)

//  Go语言中标识符由字母数字和_(下划线）组成，并且只能以字母和_开头。 举几个例子：abc, _, _123, a123。

// 函数外只能放置(变量/常量/函数/类型)的声明
var a string

func _string() {
	a = "123"
	a = "123456"
	b := "456"
	c := "str := \"c:\\Code\\lesson1\\go.exe\""
	fmt.Printf("变量:%v\n", a)
	// 转义符
	// \r	回车符（返回行首）
	// \n	换行符（直接跳到下一行的同列位置）
	// \t	制表符
	// \'	单引号
	// \"	双引号
	// \\	反斜杠
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
	fmt.Println(`
	123
	12
	`) // 多行

	// 	len(str)	求长度
	// +或fmt.Sprintf	拼接字符串
	// strings.Split	分割
	// strings.Contains	判断是否包含
	// strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
	// strings.Index(),strings.LastIndex()	子串出现的位置
	// strings.Join(a[]string, sep string)	join操作
	fmt.Println(len(a))
	fmt.Println(fmt.Sprintf("%s213%s", a, b))
	ret := strings.Split(c, "\\")
	fmt.Printf("%T\n", ret)
	fmt.Println(ret)
	fmt.Println(strings.Contains(c, "\\"))
	fmt.Println(strings.HasPrefix(c, "str"))
	fmt.Println(strings.Index(c, "str"))
	fmt.Println(strings.Join(ret, "+"))
}

func _btye() {
	// 字符用单引号（’）包裹起来，如：a := '中'
	// 	uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
	// rune类型，代表一个 UTF-8字符。
	// 当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32
	d := '中'
	fmt.Println(d)

}

// Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。
// 遍历字符串
func traversalString() {
	s := "hello沙河사샤"
	for i := 0; i < len(s); i++ { // s[i]: byte
		fmt.Printf("%v(%c) %T %T ", s[i], s[i], s[i], i) // %c:字符
	}
	fmt.Println()
	for _, r := range s { // rune eg: 'h' 'e'...
		fmt.Printf("(%c) %T ", r, r)
		fmt.Println(r)
	}
	fmt.Println()
}

// 修字符串
func changeString() {
	// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
	d := "abczAZ??"
	f := d + "汉字的数量"
	rune32 := []rune(f)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
	fmt.Println(rune32)
	fmt.Printf("%T\n%c\n", rune32, rune32)
	var length int = 0
	for _, r := range rune32 {
		fmt.Printf("%c %v\n", r, r)
		if r > 122 {
			length = length + 1
		}
	}
	fmt.Println(length)
}

// 整型
/*
 按长度分为有符号整型：int8、int16、int32、int64 对应的无符号整型：uint8、uint16、uint32、uint64
 其中，uint8就是我们熟知的byte型，int16对应C语言中的short型，int64对应C语言中的long型。
*/

func _int() {
	var n int8
	// n = 1111111111 // overflows
	n = 127
	v := 1_000 // _: 分隔数字 方便观看
	// int 32位操作系统上就是int32，64位操作系统上就是int64
	fmt.Println(n)
	fmt.Println(v)

	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	x := a
	a++
	fmt.Println(x)
	fmt.Println(a)
}

// 入口函数
func main() {
	fmt.Println("字符串:")
	_string()
	fmt.Println()
	fmt.Println()
	fmt.Println("字符:")
	_btye()
	traversalString()
	changeString()
	fmt.Println()
	fmt.Println("整型:")
	_int()
}
