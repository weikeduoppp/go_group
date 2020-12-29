package logger

import "fmt"

const (
	DEBUG = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// Logger 日志结构体
type Logger struct {
}

// NewLog Logger 构造函数
func NewLog() Logger {
	return Logger{}
}

// Debug 调试
func (l Logger) Debug(msg interface{}) {
	fmt.Println(msg)
}
