package logger

import (
	"fmt"
	"time"
)

// DEBUG... level
const (
	DEBUG LogLevel = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// LogLevel log级别
type LogLevel uint8

// Logger 日志结构体
type Logger struct {
	level []LogLevel
}

// NewLog Logger 构造函数
func NewLog() Logger {
	return Logger{}
}

// checkLevel 检测是否包含
func (l Logger) checkLevel(level LogLevel) (b bool) {
	for _, v := range l.level {
		if v == level {
			b = true
		}
	}
	return
}

/*
	获取行号信息:
	runtime.Caller func Caller(skip int) (pc uintptr, file string, line int, ok bool)
	Caller报告当前go程调用栈所执行的函数的文件和行号信息。实参skip为上溯的栈帧数，0表示Caller的调用者（Caller所在的调用栈）。（由于历史原因，skip的意思在Caller和Callers中并不相同。）函数的返回值为调用栈标识符、文件名、该调用在文件中的行号。如果无法获得信息，ok会被设为false。

	获取函数名: runtime.FuncForPC(pc) FuncForPC返回一个表示调用栈标识符pc对应的调用栈的*Func；如果该调用栈标识符没有对应的调用栈，函数会返回nil。每一个调用栈必然是对某个函数的调用。
	获取路径名: path.Base(path string)
*/
// pln 公用
func (l Logger) pln(msg interface{}, level LogLevel) {
	// 过滤
	if !l.checkLevel(level) {
		return
	}
	now := time.Now()
	var levelInfo string
	switch level {
	case DEBUG:
		levelInfo = "DEBUG"
	case TRACE:
		levelInfo = "TRACE"
	case INFO:
		levelInfo = "INFO"
	case WARNING:
		levelInfo = "WARNING"
	case ERROR:
		levelInfo = "ERROR"
	case FATAL:
		levelInfo = "FATAL"
	default:
		levelInfo = "unknown"
	}
	fmt.Printf("[%s] [%s] %v \n", now.Format("2006-01-02 15:04:05"), levelInfo, msg)
}

// Debug 调试
func (l Logger) Debug(msg interface{}) {
	l.pln(msg, DEBUG)
}

// Trace 跟踪
func (l Logger) Trace(msg interface{}) {
	l.pln(msg, TRACE)
}
