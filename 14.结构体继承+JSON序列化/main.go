package main

// go fmt 格式化
import (
	"encoding/json"
	"fmt"
)

type animal struct {
	// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。 大写json化才会被看到.
	Name     string `json:"name" db:"name" ini:"name"` // `key1:"value1" key2:"value2"` 转化重命名
	Age      int8   `json:"age"`
	hobbies  [3]string
	hobbies2 []string
}

func (t *animal) move() {
	fmt.Printf("%v move\n", t.Name)
}

func (t *animal) setHobbies(hobbies [3]string) { // 数组ok
	t.hobbies = hobbies
}
func (t *animal) setHobbies2(hobbies []string) {
	// t.hobbies2 = hobbies // hobbies 本体修改 hobbies2就修改
	// slice和map这两种数据类型都包含了指向底层数据的指针 正确方法
	b := make([]string, len(hobbies))
	copy(b, hobbies)
	t.hobbies2 = b
}

type cat struct {
	cate string
	animal
}

func (c *cat) miao() {
	fmt.Printf("%v 喵\n", c.Name)
}

func _func() {
	c := cat{
		cate: "猫",
		animal: animal{
			Name: "安娜",
		},
	}
	c.move()
	c.miao()
	fmt.Println(c.Name)
}

func _json() {
	b := cat{
		cate: "猫", // 私有
		animal: animal{
			Name: "安娜",
			Age:  18,
		},
	}
	data, err := json.Marshal(b) // 转字节切片
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\n", data)
	str := `{"name": "秃头","age":18}`
	fmt.Printf("%T\n", str)
	c := &cat{}
	err = json.Unmarshal([]byte(str), c) // 指针类型 让json可以修改c
	if err != nil {
		fmt.Println(err)
		return
	}
	c.cate = "猫"
	c.animal.Name = "战神"
	fmt.Printf("%#v\n", c)
}

// slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意
func mark() {
	a := animal{
		Name: "牛",
	}
	b := [3]string{"吃草", "犁地"}
	c := []string{"吃草", "犁地"}
	a.setHobbies(b)
	a.setHobbies2(c)
	fmt.Printf("%#v\n", a)
	b[0] = "睡觉"
	c[0] = "睡觉"
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%#v\n", a)
}

// 入口函数
func main() {
	fmt.Println()
	fmt.Println("结构体继承:")
	_func()
	fmt.Println()
	fmt.Println("结构体与JSON序列化:")
	_json()
	fmt.Println()
	fmt.Println("slice和map这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们时要特别注意:")
	mark()
}
