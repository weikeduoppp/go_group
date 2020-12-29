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

// pln 公用
func (l Logger) pln(msg interface{}, level LogLevel) {
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
