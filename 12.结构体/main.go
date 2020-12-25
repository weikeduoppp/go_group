package main

// go fmt 格式化
import (
	"fmt"
	"unsafe"
)

type customInt int  // 自定义类型
type aliasInt = int // 类型别名

func _func() {
	var n customInt = 0
	var m aliasInt = 0
	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", m)
	// rune 是int32的别名
	var a rune = '中'
	fmt.Printf("%T\n", a)
}

// struct "类"  如果首字母大写.一般是公开的
type person struct {
	name     string
	age      uint8
	gender   string
	sayhello func(msg string)
}

func _struct() {
	var xo person
	xo.name = "xo"
	xo.age = 18
	xo.gender = "男"
	fmt.Println(xo)
	fmt.Printf("%T\n", xo)

	// 匿名结构体 临时
	var temp struct {
		name string
	}
	temp.name = "temp"
	fmt.Println(temp)
	fmt.Printf("%T\n", temp)

	// 初始化方式
	// 指针 ==  &person{}
	var A = new(person)
	(*A).name = "啊" // 指针赋值可以省略
	A.name = "啊"    // 在Go语言中支持对结构体指针直接使用.来访问结构体的成员。
	fmt.Println(A)
	fmt.Printf("%T\n", A)
	fmt.Printf("%p\n", A)  // A的地址
	fmt.Printf("%p\n", &A) // 值的地址

	// 1.键值
	B := person{
		name: "a",
		// ... 没有初始化的结构体，其成员变量都是对应其类型的零值。
	}

	// 2.填充顺序必须与字段在结构体中的声明顺序一致 也可以指针&
	C := &person{
		"C",
		18,
		"男",
		func(msg string) { fmt.Println(msg) },
	}

	fmt.Println(B)
	fmt.Println(C)

	// 结构体占用一块连续的内存。
	type test struct { // 升序
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	// 空结构体是不占用空间的。 调整结构体内成员变量的字段顺序就能达到缩小结构体占用大小
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) // 0
	// 在 Go 中恰到好处的内存对齐: https://segmentfault.com/a/1190000017527311?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com

}

// 构造函数
func newPerson(name string, age uint8, gender string) *person {
	// 当结构体体积比较大. 推荐返回结构体指针, 减少内存开销
	return &person{
		name:   name,
		age:    age,
		gender: gender,
		sayhello: func(msg string) {
			fmt.Println(msg)
		},
	}
}

// 方法和接收者
/*
	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
	}
	接收者类型: 自定义类型(结构体/type)
*/
func (d person /* 值类型的接收者(副本) */) sayNo() { // 作用于特定类型的函数. (指定哪种接收者类型能用的函数) d 类似 this, self; 小写字母开头
	fmt.Printf("%v say no\n", d.name)
	// 如果有修改操作只是针对副本，无法修改接收者变量本身。
}

// 指针类型的接收者 由于指针的特性，调用方法时修改接收者指针的任意成员变量
func (d *person) setAge(newAge uint8) {
	d.age = newAge
}

/*
	什么时候应该使用指针类型接收者
	- 需要修改接收者中的值
	- 接收者是拷贝代价比较大的大对象
	- 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/

//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}

//  非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。

// 匿名结构体 嵌套结构体(可匿名)  !!! 当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。
func anonymous() {
	type students struct {
		name string
		age  int
	}
	type teacher struct {
		name string
		age  int
	}
	type address struct {
		addr string
		size int
	}
	type class struct {
		title    string
		students // 可嵌套
		teacher
		address // 可继承
	}

	a := students{
		name: "ant",
		age:  8,
	}
	fmt.Println(a.name)
	c := class{
		title: "一班",
		students: students{
			name: "李四",
			age:  18,
		},
		teacher: teacher{
			name: "老师",
		},
		address: address{
			addr: "位置",
		},
	}
	fmt.Println(c.address.addr) // 可省略匿名字段address
	fmt.Println(c.students.name)
	fmt.Println(c.teacher.name) // 如果匿名结构体字段有冲突则不能省略
}

// 入口函数
func main() {
	fmt.Println()
	fmt.Println("自定义类型,类型别名:")
	_func()
	fmt.Println()
	fmt.Println("struct:")
	_struct()
	fmt.Println()
	fmt.Println("struct-构造函数:")
	p := newPerson("dog", 18, "公")
	fmt.Println(unsafe.Sizeof(p)) // 8 如果不是指针的话, 则完整person空间 eg: 40
	fmt.Println(p.name)
	p.sayhello("hello, dog")
	p.sayNo()
	fmt.Println(p.age)
	p.setAge(22)
	fmt.Println(p.age)
	fmt.Println()
	fmt.Println("任意类型添加方法:")
	var m1 MyInt
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
	fmt.Println()
	fmt.Println("匿名字段(取类型):")
	anonymous()
}
