package main

import (
	"fmt"
)

/*
type 接口 interface {
	方法(参数...) (返回值...)
	...
}
*/

// 接口一般以er后缀 一种特殊的类型 只关心方法, 约束有某种方法
type sayer interface {
	say()
}
type mover interface {
	move()
}

// 接口嵌套
type animaler interface {
	sayer
	mover
}

type animal struct {
	Name string `json:"name"`
}

func (a animal) move() {
	fmt.Printf("%v 在动\n", a.Name)
}

type cat struct {
	*animal
}
type dog struct {
	*animal
}

func (c cat) say() { // 如果是值接收者 变量可以值类型. 也可以指针类型
	fmt.Println(c.Name)
	fmt.Println("喵~")
}
func (d *dog) say() { // 如果实现sayer接口的是*dog类型
	fmt.Println(d.Name)
	fmt.Println("汪汪~")
}

// 一个类型实现多个接口
// 并且一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。
// 接口嵌套

func _func() {
	var x sayer // 变量
	var a mover
	var c = cat{
		animal: &animal{
			Name: "安娜",
		},
	}
	var d = &dog{ //  那传入sayer也需要是*dog类型
		animal: &animal{
			Name: "旺财",
		},
	}
	x = &c // c实现了sayer接口的中所有的方法 所以可以赋值x
	c.say()
	x.say()
	x = d
	x.say()
	a = d
	a.move()
}

// interface{}
// 空接口
func _func2() {
	// 空接口类型的变量可以存储任意类型的变量。

	// 使用空接口实现可以接收任意类型的函数参数。
	open := func(a interface{}) {
		fmt.Println(a)
	}
	open("aa")
	// 空接口作为map的值
	teacher := make(map[string]interface{}, 10)
	teacher["name"] = "老师"
	teacher["age"] = 18
	fmt.Println(teacher)

}

// 空接口可以存储任意类型的值，那我们如何获取其存储的具体数据呢？
func assign(a interface{}) {
	str, ok := a.(string)
	if !ok {
		fmt.Println("a is not string")
	} else {
		fmt.Println(str)
	}
	switch t := a.(type) {
	case string:
		fmt.Println("a is string")
	default:
		fmt.Println(t)
	}
}

// 入口函数
func main() {
	fmt.Println()
	fmt.Println("接口:")
	_func()
	fmt.Println()
	fmt.Println("空接口:")
	_func2()
	fmt.Println()
	fmt.Println("类型断言 x.(T):")
	assign("a")
	assign(100)
}
