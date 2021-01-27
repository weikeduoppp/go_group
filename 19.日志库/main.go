package main

import (
	"fmt"
	"log"
	"logger"
	"time"
)

// Log 全局log
var Log logger.Console

// FileLog ...
var FileLog *logger.FileLog

func test() {
	// 允许的日志级别
	level := []logger.LogLevel{logger.DEBUG, logger.INFO, logger.TRACE, logger.WARNING, logger.ERROR}
	Log = logger.NewLog(level)
	Log.Debug("console")
	Log.Fatal("console")
	Log.Info("console")
	Log.Info("console %v", Log)
}

// FileTest ...
func FileTest() {
	// 允许的日志级别
	level := []logger.LogLevel{logger.DEBUG, logger.INFO, logger.TRACE, logger.WARNING, logger.ERROR, logger.FATAL}
	FileLog = logger.NewFileLog(level, "./logs", 1*1024)
	for {
		FileLog.Debug("console")
		FileLog.Fatal("console")
		FileLog.Info("console")
		FileLog.Info("console %v", FileLog)
		time.Sleep(time.Millisecond * 500)
	}
}

func _func() {
	// 设置输出位置 默认终端
	// log.SetOutput(os.Stdout)
	// log 简单的日志服务
	log.Println("输出")
}

func main() {
	fmt.Println()
	fmt.Println("log:")
	_func()
	fmt.Println()
	fmt.Println("测试日志logger:")
	test()
	fmt.Println()
	fmt.Println("测试日志FileTest:")
	FileTest()
}
