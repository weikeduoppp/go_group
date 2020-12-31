package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

// DEBUG... level
const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// LogLevel log级别
type LogLevel uint8

/*
	获取行号信息:
	runtime.Caller func Caller(skip int) (pc uintptr, file string, line int, ok bool)
	Caller报告当前go程调用栈所执行的函数的文件和行号信息。实参skip为上溯的栈帧数，0表示Caller的调用者（Caller所在的调用栈）。（由于历史原因，skip的意思在Caller和Callers中并不相同。）函数的返回值为调用栈标识符、文件名、该调用在文件中的行号。如果无法获得信息，ok会被设为false。

	获取函数名: runtime.FuncForPC(pc) FuncForPC返回一个表示调用栈标识符pc对应的调用栈的*Func；如果该调用栈标识符没有对应的调用栈，函数会返回nil。每一个调用栈必然是对某个函数的调用。
	获取路径名: path.Base(path string)
*/
func getLineInfo(skip int) (funcName, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime caller failed")
		return
	}
	funcName = strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]
	fileName = path.Base(file)
	return
}

func getLevelInfo(level LogLevel) (levelInfo string) {
	switch level {
	case DEBUG:
		levelInfo = "DEBUG  "
	case TRACE:
		levelInfo = "TRACE  "
	case INFO:
		levelInfo = "INFO   "
	case WARNING:
		levelInfo = "WARNING"
	case ERROR:
		levelInfo = "ERROR  "
	case FATAL:
		levelInfo = "FATAL  "
	default:
		levelInfo = "UNKNOWN"
	}
	return
}

func getFileInfo(file *os.File) (name string, size int64, modTime time.Time) {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Get file info failed err: %v\n", err)
	}
	return fileInfo.Name(), fileInfo.Size(), fileInfo.ModTime()
}
