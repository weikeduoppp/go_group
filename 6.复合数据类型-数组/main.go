package main

// go fmt 格式化
import (
	"fmt"
)

func _array() {
	// 声明必须指定长度 类型
	// 数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。 [5]int和[10]int是不同的类型。
	var (
		a [3]uint
		b [4]uint
	)
	// a = b // 不同类型
	fmt.Println(a, b)

	// 初始化 (布尔值默认 false | uint, int, float: 0 | string: "")
	a1 := [2]int{1, 2}
	fmt.Println(a1)

	a2 := [...]string{"123", "213"} // 自动推断长度
	fmt.Println(a2)
	fmt.Println(len(a2))

	a3 := [3]string{0: "上海", 2: "深圳"} // 索引初始化.
	fmt.Println(a3)
	fmt.Println(a3[1] == "") // ""

	// 索引遍历
	for i := 0; i < len(a2); i++ {
		fmt.Println(a2[i])
	}

	// range遍历
	for i, v := range a3 {
		fmt.Println(i, v)
	}

	// 多维数组  多维数组只有第一层可以使用...来让编译器推导数组长度
	// [[1,2,3],[2,3], [4,5]]
	a4 := [...][2]int{[2]int{2, 3}, [2]int{1, 3}}
	fmt.Printf("%T %v\n", a4, a4)
}

// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。!!
func modifyArray(x [3]int) {
	x[0] = 100
	fmt.Println(x)
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
	fmt.Println(x)
}

func _find() {
	a := [5]int{1, 3, 5, 7, 8}
	for i, j := range a {
		for l, k := range a[i+1:] {
			if j+k == 8 {
				fmt.Printf("(%d, %d)  ", i, l+i+1)
			}
		}
	}
}

func test() {
	a := [...]uint{2, 1, 3, 5, 4, 4, 7, 8, 7}
	// 找出和为8的两个下标
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				fmt.Printf("(%v,%v)", i, j)
			}
		}
	}


// 入口函数
func main() {

	fmt.Println()
	fmt.Println("数组:")
	_array()
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
	test()
	// _find()
}
