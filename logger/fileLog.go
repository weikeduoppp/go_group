package logger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLog 日志结构体
type FileLog struct {
	level       []LogLevel //
	filePath    string     // logs目录
	logFile     *os.File   // 正常log
	errFile     *os.File   // 错误log
	maxFileSize int64      // 文件大小 大于则切割
}

// func (f *File) Stat() (fi FileInfo, err error) // Stat返回描述文件f的FileInfo类型值。如果出错，错误底层类型是*PathError。
/*
	type FileInfo interface {
    Name() string       // 文件的名字（不含扩展名）
    Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
    Mode() FileMode     // 文件的模式位
    ModTime() time.Time // 文件的修改时间
    IsDir() bool        // 等价于Mode().IsDir()
    Sys() interface{}   // 底层数据来源（可以返回nil）
	}
*/

// NewFileLog FileLog 构造函数
func NewFileLog(level []LogLevel, filePath string, maxFileSize int64) *FileLog {
	file := &FileLog{
		level:       level,
		filePath:    filePath,
		maxFileSize: maxFileSize,
	}
	err := file.initFile(true, true) // 打开日志文件
	if err != nil {
		panic(err)
	}
	return file
}

// initFile 生成logs/...
func (f *FileLog) initFile(logBool, errBool bool) error {
	if logBool {
		FilePath := f.filePath + "/out.log"
		file, err := os.OpenFile(FilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Printf("%v open failed err: %v\n", FilePath, err)
			return err
		}
		f.logFile = file
	}
	if errBool {
		errFilePath := f.filePath + "/err.log"
		errFile, err := os.OpenFile(errFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Printf("%v open failed err: %v\n", errFilePath, err)
			return err
		}
		f.errFile = errFile
	}
	return nil
}

// checkLevel 检测是否包含
func (f *FileLog) checkLevel(level LogLevel) (b bool) {
	for _, v := range f.level {
		if v == level {
			b = true
		}
	}
	return
}

// 检测
func (f *FileLog) checkFileSizeAndSplitAndBackup(file *os.File, logBool, errBool bool) error {
	name, size, modTime := getFileInfo(file)
	if size >= f.maxFileSize {
		/*
			1. 检测超出大小关闭文件
			2. 重命名(备份)
				file_backup_20060102150405
			3. 重新赋值
		*/
		file.Close()
		oldpath := path.Join(f.filePath, name)
		newName := fmt.Sprintf("backup_%s_%s", modTime.Format("20060102150405"), name)
		newpath := path.Join(f.filePath, newName)
		os.Rename(oldpath, newpath)
		err := f.initFile(logBool, errBool)
		return err
	}
	return nil
}

// pln 写入
func (f *FileLog) pln(level LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	// 过滤
	if !f.checkLevel(level) {
		return
	}
	now := time.Now()
	levelInfo := getLevelInfo(level)

	funcName, fileName, line := getLineInfo(3)

	// 正常log
	if level < ERROR {
		// 大于则切割
		err := f.checkFileSizeAndSplitAndBackup(f.logFile, true, false)
		if err != nil {
			fmt.Printf("logFile split file Backup failed err: %v\n, msg: %v", err, msg)
			return
		}
		fmt.Fprintf(f.logFile, "[%s] |%s| [%s line: %d >> %s] %v \n", now.Format("2006-01-02 15:04:05"), levelInfo, fileName, line, funcName, msg)
	} else { // errLog
		err := f.checkFileSizeAndSplitAndBackup(f.errFile, false, true)
		if err != nil {
			fmt.Printf("logFile split file Backup failed err: %v\n, msg: %v", err, msg)
			return
		}
		fmt.Fprintf(f.errFile, "[%s] |%s| [%s line: %d >> %s] %v \n", now.Format("2006-01-02 15:04:05"), levelInfo, fileName, line, funcName, msg)
	}
}

// Debug 调试
func (f *FileLog) Debug(format string, a ...interface{}) {
	f.pln(DEBUG, format, a...)
}

// Trace 跟踪
func (f *FileLog) Trace(format string, a ...interface{}) {
	f.pln(TRACE, format, a...)
}

// Info 消息
func (f *FileLog) Info(format string, a ...interface{}) {
	f.pln(INFO, format, a...)
}

// Warning 警告
func (f *FileLog) Warning(format string, a ...interface{}) {
	f.pln(WARNING, format, a...)
}

// Error 错误
func (f *FileLog) Error(format string, a ...interface{}) {
	f.pln(ERROR, format, a...)
}

// Fatal 程序未找到
func (f *FileLog) Fatal(format string, a ...interface{}) {
	f.pln(FATAL, format, a...)
}
