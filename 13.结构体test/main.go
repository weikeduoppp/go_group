package main

// go fmt 格式化
import (
	"fmt"
)

type student struct {
	id   int64
	name string
}

type class struct {
	allStudent map[int64]*student
}

func (c class) addStudent() {
	var (
		name string
		id   int64
	)
	fmt.Println("输入学号和名字, 空格隔开")
	fmt.Scanln(&id, &name)
	c.allStudent[id] = &student{
		id:   id,
		name: name,
	}
}
func (c class) deleteStudent(id int64) {
	delete(c.allStudent, id)
}
func (c class) selectAll() {
	for _, v := range c.allStudent {
		// n位数. 缺失补充0   %0nd
		fmt.Printf("学号: %03d, 姓名: %v \n", v.id, v.name)
	}
}

func _func() {
	A := class{
		allStudent: make(map[int64]*student, 50),
	}
	fmt.Println("A班学生表")
	A.selectAll()
	A.addStudent()
	A.addStudent()
	A.addStudent()
	fmt.Println("添加后: A班学生表")
	A.selectAll()
	A.deleteStudent(1)
	fmt.Println("删除后: A班学生表")
	A.selectAll()
}

// 入口函数
func main() {
	fmt.Println()
	fmt.Println("生成一个班的管理系统:")
	_func()
}
