package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"17.文件操作os/write"
)

// 在Go语言程序执行时导入包语句会自动触发包内部init()函数的调用。需要注意的是： init()函数没有参数也没有返回值。 init()函数在程序运行时自动被调用执行，不能在代码中主动调用它。

func init() {
	fmt.Println()
	fmt.Println("指定读取多少:")
	_func()
	fmt.Println()
	fmt.Println("指定读取一行:")
	_readFileLine()
	fmt.Println()
	fmt.Println("指定读取整个文件:")
	_readFile()
}

// 指定读取多少
func _func() {
	// func Open(name string) (file *File, err error)
	fmt.Println("os")
	file, err := os.Open("./init.go")
	if err != nil {
		fmt.Printf("open failed %v\n", err)
		return
	}
	// 记得关闭文件
	defer file.Close()

	var content []byte
	var temp = make([]byte, 128)
	// func (f *File) Read(b []byte) (n int, err error)
	for {
		n, err := file.Read(temp[:])
		// 读到文件末尾时会返回0和io.EOF
		if err == io.EOF {
			fmt.Printf("已读取完 %v\n", err)
			break
		}
		if err != nil {
			fmt.Printf("read failed %v\n", err)
			return
		}
		content = append(content, temp[:n]...)
	}
	fmt.Println(string(content))
}

// bufio是在file的基础上封装了一层API，支持更多的功能。 缓存中操作
func _readFileLine() {
	file, err := os.Open("./init.go")
	if err != nil {
		fmt.Printf("open failed %v\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("已读取完 %v\n", err)
			break
		}
		if err != nil {
			fmt.Printf("open failed %v\n", err)
			return
		}
		fmt.Print(line)
	}
}

// io/ioutil包的ReadFile方法能够读取完整的文件，只需要将文件名作为参数传入。
func _readFile() {
	content, err := ioutil.ReadFile("./go.mod")
	if err != nil {
		fmt.Printf("read file failed, err: %v", err)
		return
	}
	fmt.Println(string(content))
}

func _write() {
	write.MyWriteFile()
}

// mycat 模拟linux cat指令
func mycat() {
	var path string
	fmt.Scanf("cat %s", &path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("read file failed, err: %v", err)
		return
	}
	fmt.Println(string(content))
}

// useBufio 输入带空格的 eg:a b c
func useBufio() {
	reader := bufio.NewReader(os.Stdin) // 标准输入
	fmt.Println("请输入:")
	s, _ := reader.ReadString('\n')
	fmt.Printf("你输入的是: %v\n", s)
}

// cat命令实现
func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n') //注意是字符
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func useCat() {
	flag.Parse() // 解析命令行参数
	if flag.NArg() == 0 {
		// 如果没有参数默认从标准输入读取内容
		cat(bufio.NewReader(os.Stdin))
	}
	// 依次读取每个指定文件的内容并打印到终端
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}

func main() {
	fmt.Println()
	fmt.Println("文件写入")
	_write()
	fmt.Println()
	fmt.Println("模拟cat")
	// mycat() // 自己
	fmt.Println()
	fmt.Println("useBufio")
	// useCat() // 教程
	useBufio()
}
