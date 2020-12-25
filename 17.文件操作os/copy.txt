package write

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。
/*
	func OpenFile(name string, flag int, perm FileMode) (*File, error) {
		...
	}

	os.O_WRONLY	只写
	os.O_CREATE	创建文件
	os.O_RDONLY	只读
	os.O_RDWR	读写
	os.O_TRUNC	清空
	os.O_APPEND	追加
*/
func writeFileDemo1() {
	file, err := os.OpenFile("xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("write failed err: %v", err)
		return
	}
	defer file.Close()
	str := "hello ye"
	file.Write([]byte(str))
	file.WriteString("!!!!!")
}

// bufio.NewWriter
func writeFileDemo2() {
	file, err := os.OpenFile("xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("write failed err: %v", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("存入缓存写入")
	// 写入
	writer.Flush()
}

// ioutil.WriteFile
func writeFileDemo3() {
	str := "ioutil.WriteFile"
	err := ioutil.WriteFile("xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("write failed err: %v", err)
		return
	}
}

// CopyFile io.Copy()
func CopyFile(targetName string, srcName string) (written int64, err error) {
	// 1.打开文件 2.创建文件 3.复制
	file, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open failed %v\n", err)
		return
	}
	defer file.Close()
	targetFile, err := os.OpenFile(targetName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("create failed %v\n", err)
		return
	}
	defer targetFile.Close()
	return io.Copy(targetFile, file)
}

// MyWriteFile 文件写入功能
func MyWriteFile() {
	// writeFileDemo1()
	// writeFileDemo2()
	writeFileDemo3()
	fmt.Println("文件复制")
	CopyFile("copy.txt", "./write/write.go")
}
