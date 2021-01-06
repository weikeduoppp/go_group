package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// MysqlConf 数据库
type MysqlConf struct {
	Hostname string `ini:"hostname"`
	Port     int    `ini:"port"`
	Username string `ini:"usename"`
	Password string `ini:"password"`
}

// RedisConf ...
type RedisConf struct {
	Hostname string `ini:"hostname"`
	Port     int    `ini:"port"`
}

// Conf ...
type Conf struct {
	MysqlConf `ini:"mysql"`
	RedisConf `ini:"redis"`
}

// IniUnMarshal ini读取
func IniUnMarshal(filename string, d interface{}) (err error) {
	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	// 0 参数类型校验
	if t.Kind() != reflect.Ptr {
		err = fmt.Errorf("data type is not reflect.Ptr")
	}
	// 0.1 d => 指针类型 且是结构体
	if v.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("data type is not reflect.Struct")
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	lines := strings.Split(string(content), "\n")
	var sectionName string
	for idx, line := range lines {
		// 1 按行读取文件
		// 1.1 注释&&空行 跳过 去除首尾空格  # ;
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}

		// 1.2 [T] 获取节 首尾对应[] 防止内容空值. ==> tag
		if strings.HasPrefix(line, "[") {
			if !strings.HasPrefix(line, "[") || !strings.HasSuffix(line, "]") {
				err = fmt.Errorf("line: %d, syntax error", idx+1)
				return
			}
			section := strings.TrimSpace(line[1 : len(line)-1])
			if len(section) != 0 {
				// 找到结构体
				for i := 0; i < t.Elem().NumField(); i++ {
					field := t.Elem().Field(i)
					if field.Tag.Get("ini") == section {
						sectionName = field.Name
					}
				}
			} else {
				err = fmt.Errorf("line: %d, syntax error", idx+1)
				return
			}
		} else {
			if strings.HasSuffix(line, "]") {
				err = fmt.Errorf("line: %d, syntax error", idx+1)
				return
			}
			// 不包含
			if !strings.Contains(line, "=") {
				err = fmt.Errorf("line: %d, syntax error", idx+1)
				return
			}
			// 1.3 key=value 获取键值对 根据=切割
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			if len(key) == 0 {
				err = fmt.Errorf("line: %d, syntax error", idx+1)
				return
			}
			fieldValue := v.Elem().FieldByName(sectionName)
			fieldValueType := fieldValue.Type()
			if fieldValue.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", sectionName)
				return
			}
			// 寻找对应字段
			var fieldName string
			var fieldType reflect.Kind
			for i := 0; i < fieldValue.NumField(); i++ {
				field := fieldValueType.Field(i)
				if field.Tag.Get("ini") == key {
					fmt.Printf("找到%v字段, type is %v\n", key, field.Type)
					fieldName = field.Name
					fieldType = field.Type.Kind()
					break
				}
			}
			// fieldValue里找不到该字段
			if len(fieldName) == 0 {
				continue
			}
			currentField := fieldValue.FieldByName(fieldName)
			switch fieldType {
			case reflect.String:
				currentField.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				currentField.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				currentField.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				currentField.SetFloat(valueFloat)
			}
			// 检测到节 获取对应的结构体 分配赋值   v.FieldByName() v需要是结构体
		}

	}
	return
}

func main() {
	var conf Conf
	err := IniUnMarshal("./conf.ini", &conf)
	if err != nil {
		fmt.Printf("load ini failed, err: %v", err)
	}
	fmt.Printf("%#v\n", conf)
}
