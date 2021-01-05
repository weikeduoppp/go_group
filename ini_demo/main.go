package main

// MysqlConf 数据库
type MysqlConf struct {
	Hostname string `ini:"hostname"`
	Port     int    `ini:"port"`
	Username string `ini:"usename"`
	Password string `ini:"password"`
}

// Conf ...
type Conf struct {
	MysqlConf
}

// IniUnMarshal ini读取
func IniUnMarshal(file string, d interface{}) (err error) {
	// 0 参数类型校验
	// 0.1 d => 指针类型 且是结构体
	// 1 按行读取文件
	// 1.1 注释跳过
	// 1.2 [T] 获取节
	// 1.3 key=value 获取键值对 根据=切割
	// 检测到节 分配赋值
	return
}

func main() {

}
