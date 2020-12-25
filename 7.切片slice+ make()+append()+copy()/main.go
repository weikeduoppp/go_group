package main

// go fmt 格式化
import (
	"fmt"
	"sort"
)

func _slice() {
	// !! 切片是引用类型  []T
	// var s1 []int
	var s2 []string
	// 初始化(底层数组)  自动扩容 容量: cap()
	s1 := []int{12, 3}
	fmt.Println(s1)
	fmt.Println(s2 == nil) // nil 空
	fmt.Println(cap(s1))   // 容量

	// 由数组得到切片
	a := [5]int{1, 3, 5, 7, 8}
	s3 := a[0:3]                                                    // [1, 3, 5]  [x:n]: 从x(不填为0)下标开始取到n(不包括n, 不填取到尾部)下标为止的切片 [0,3}
	fmt.Printf("%v  %T, 长度: %v 容量: %v\n", s3, s3, len(s3), cap(s3)) // 底层是从数组切出来的. 容量从开始下标开始算到数组尾部
	// 切片再切片
	s4 := s3[:1]
	fmt.Printf("%v , 容量: %v\n", s4, cap(s4))

	// 证明切片是引用类型 指向底层数组 修改的是底层数组
	a[1] = 1000
	fmt.Println(s3, s4)
	s3[0] = 22
	fmt.Println(s3, s4, a)

}

func _make() {
	// 使用make()函数创建切片 make([]T, size, cap) 已初始化
	a1 := make([]int, 3, 5)
	fmt.Printf("%v  %T, 长度: %v 容量: %v\n", a1, a1, len(a1), cap(a1))
	// 证明切片长度为0. 要用len(T) == 0来判断.不是 T == nil
	// var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	// s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	// s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
}

func _append() {
	// Go语言的内建函数append()可以为切片动态添加元素。 可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（x,T…）。
	var s []int
	s = append(s, 1) // 同时初始化
	s = append(s, 2, 3)
	s2 := []int{4, 5, 6}
	s = append(s, s2...)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	// 每次扩容后都是扩容前的2倍
	// $GOROOT/src/runtime/slice.go源码
	/*
		- 首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
		- 否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
		- 否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
		- 如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
	*/
}

func _copy() {
	// copy(destSlice, srcSlice []T) copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中, 前提目的切片初始化过且有长度.
	a := []int{1, 2, 3}
	b := make([]int, 3, 3)
	fmt.Println(b)
	copy(b, a)
	fmt.Println(b)
	b[0] = 111
	// 修改不影响复制源
	fmt.Println(a)
	fmt.Println(b)
}

// 通过append从切片中删除元素
func _delete(index int) {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:index], a[index+1:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}

func _test() {
	var a = make([]string, 5, 10)
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}
func _test2() {
	var a = make([]int, 5, 10)
	for i := 10; i > 0; i-- {
		a = append(a, i)
	}
	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)

}

// 指针 GO: 指针只读
func _pointer() {
	// 1. &:取指针地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p) // *int：int类型的指针
	// 2. *：取指针地址的值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m) // int

	var a1 *int // nil pointer
	fmt.Println(a1)
	var a2 = new(int) // new函数申请一个内存地址 return 内存地址
	fmt.Println(a2)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)
}

// fmt.Sprintf("%v", T) 数字转字符串

/*
	new和make的区别
	都是开辟空间. 申请内存.
	new是给基础类型.或者结构体struct. 返回: 对应变量的指针地址
	make是给特定类型. slice[] map[T1]T2 ..  返回: 对应变量
*/

// 入口函数
func main() {

	fmt.Println()
	fmt.Println("slice:")
	_slice()
	fmt.Println()
	fmt.Println("make:")
	_make()
	fmt.Println()
	fmt.Println("append:")
	_append()
	fmt.Println()
	fmt.Println("copy:")
	_copy()
	fmt.Println()
	fmt.Println("delete:")
	_delete(0)
	fmt.Println()
	fmt.Println("test:")
	_test()
	fmt.Println()
	fmt.Println("test2:")
	_test2()
	fmt.Println()
	fmt.Println("指针pointer:")
	_pointer()
}
