package logger

import (
	"fmt"
	"time"
)

// Console 日志结构体
type Console struct {
	level []LogLevel
}

// NewLog Console 构造函数
func NewLog(level []LogLevel) Console {
	return Console{
		level: level,
	}
}

// checkLevel 检测是否包含
func (l Console) checkLevel(level LogLevel) (b bool) {
	for _, v := range l.level {
		if v == level {
			b = true
		}
	}
	return
}

// pln 打印
func (l Console) pln(level LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	// 过滤
	if !l.checkLevel(level) {
		return
	}
	now := time.Now()
	levelInfo := getLevelInfo(level)
	funcName, fileName, line := getLineInfo(3)
	fmt.Printf("[%s] |%s| [%s line: %d >> %s] %v \n", now.Format("2006-01-02 15:04:05"), levelInfo, fileName, line, funcName, msg)
}

// Debug 调试
func (l Console) Debug(format string, a ...interface{}) {
	l.pln(DEBUG, format, a...)
}

// Trace 跟踪
func (l Console) Trace(format string, a ...interface{}) {
	l.pln(TRACE, format, a...)
}

// Info 消息
func (l Console) Info(format string, a ...interface{}) {
	l.pln(INFO, format, a...)
}

// Warning 警告
func (l Console) Warning(format string, a ...interface{}) {
	l.pln(WARNING, format, a...)
}

// Error 错误
func (l Console) Error(format string, a ...interface{}) {
	l.pln(ERROR, format, a...)
}

// Fatal 程序未找到
func (l Console) Fatal(format string, a ...interface{}) {
	l.pln(FATAL, format, a...)
}
