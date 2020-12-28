package write

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// os.OpenFile()函数能够以指定模式打开文件，从而实现文件写入相关功能。  FileMode：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。 linux下的权限
/*
	func OpenFile(name string, flag int, perm FileMode) (*File, error) {
		...
	}

	os.O_WRONLY	只写
	os.O_CREATE	创建文件
	os.O_RDONLY	只读
	os.O_RDWR	读写
	os.O_TRUNC	清空
	os.O_APPEND	追加 文件尾部添加
*/
func writeFileDemo1() {
	file, err := os.OpenFile("xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
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
	// 创建缓存区
	writer := bufio.NewWriter(file)
	// 写入缓存区
	writer.WriteString("存入缓存写入")
	// 写入磁盘
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
	CopyFile("write/copy.txt", "./write/write.go")
	// 编辑文件  需要一个临时文件. 再重命名 os.rename
}
