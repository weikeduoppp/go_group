package main

import (
	"fmt"
	"reflect"
)

/*
	TypeOf() ValueOf()
*/
func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
	t := reflect.TypeOf(x)
	// type name和type kind
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

type myInt int64

func _func() {
	var a = 100
	var c rune
	var b myInt
	type person struct {
		name string
		age  int
	}
	var d = person{
		name: "沙河小王子",
		age:  18,
	}

	var f = [2]int{1, 2}

	reflectType(&a)
	reflectType(c)
	reflectType(b)
	reflectType(d)
	reflectType(f) // Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func _func2() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
func _func3() {
	var a int64 = 100
	// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
}
func _func4() {
	// *int类型空指针
	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
}

func _func5() {
	type student struct {
		Name  string `json:"name"`
		Score int    `json:"score"`
	}
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
func main() {
	fmt.Println()
	fmt.Println("reflectType:")
	_func()
	fmt.Println()
	fmt.Println("reflectValue:")
	_func2()
	fmt.Println()
	fmt.Println("reflectSetValue:")
	_func3()
	fmt.Println()
	fmt.Println("IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。:")
	_func4()
	fmt.Println()
	fmt.Println("结构体反射:")
	_func5()
}
