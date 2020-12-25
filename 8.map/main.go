package main

// go fmt 格式化
import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func _map() {
	// var _nil map[string][int] // 未初始化 == nil
	// map[KeyType]ValueType
	a := map[string]string{
		"key":  "value",
		"key2": "valu2e",
	}
	fmt.Println(a["key"])
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := a["key3"]
	fmt.Println(v, ok, a["key3"] == "") // 字符串的零值: ""
	if !ok {
		fmt.Println("key3不存在")
	}
	// make(map[KeyType]ValueType, [cap]) 一般事先估算容量.避免运行期间再动态扩容
	b := make(map[string]string, 10)
	b["a"] = "b"
	b["a2"] = "b2"
	b["a3"] = "b2"

	for k, v := range b {
		fmt.Println("键值: ", k, "值: ", v)
	}

	// delete(map, key)内置删除函数
	delete(b, "a")
	fmt.Println(b)
}

// 按照指定顺序遍历map
func _sort() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

func makeMapslice() {
	// 元素类型为map类型的切片
	s1 := make([]map[int]int, 10, 10)
	s1[0] = make(map[int]int, 2) // 需要初始化
	s1[0][0] = 111
	fmt.Println(s1[0])

	// 值为切片的map
	s2 := map[string][]int{
		"key": {1, 3},
	}
	fmt.Println(s2)
}

// find字符出现次数
func findCodeFrequency() {
	str := "how do you do"
	arr := strings.Split(str, " ")
	m := make(map[string]int, 10)
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v]++
		}
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// 回文判断
func palindromeJudgment() {
	str := "32去232"
	arr := strings.Split(str, "")
	fmt.Println(arr, len(arr))
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			fmt.Println("不是回文")
			break
		}
	}
}

func watch() {
	// 切片是引用类型
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s) // [1,2,3]
	m["q1mi"] = s          // 切片
	fmt.Printf("%+v\n", m["q1mi"])
	s = append(s[:2], s[3:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}

// 入口函数
func main() {

	fmt.Println()
	fmt.Println("map:")
	_map()
	fmt.Println("sort:")
	_sort()
	fmt.Println("makeMapslice:")
	makeMapslice()
	fmt.Println()
	fmt.Println("watch:")
	watch()
	fmt.Println()
	fmt.Println("findCodeFrequency:")
	findCodeFrequency()
	fmt.Println()
	fmt.Println("palindromeJudgment:")
	palindromeJudgment()
}
