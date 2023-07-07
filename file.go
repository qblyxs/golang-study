package main

import (
	"fmt"
	"github.com/qblyxs/golang-study/filefunc"
	"os"
	"io/ioutil"
	"bufio"
	"io"
)

func main() {
	fmt.Println("==========打开文件==========")
	filefunc.SayHello()

	file15, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("打开文件 example.txt 失败")
	}
	fmt.Printf("file type: %T\n", file15)
	defer file15.Close()
	fmt.Println("==========读取文件==========")
	// 读取文件
	fmt.Println("--------一次性读取文件--------")
	// 一次性读取文件 适合小文件
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("读取文件 example.txt 失败")
	}
	fmt.Printf("文件内容: %s\n", content)


	fmt.Println("--------带缓冲区读取文件--------")
	// 带缓冲区读取文件 适合大文件
	file,err := os.Open("example.txt") // 打开文件
	if err != nil {
		fmt.Println("打开文件 example.txt 失败")
	}
	// 当函数退出时,关闭文件,释放资源
	defer file.Close()
	// 创建一个流式读取器
	reader := bufio.NewReader(file)
	// 循环读取文件内容
	for {
		// 读取到一个换行符就结束读取
		content, err := reader.ReadString('\n')
		if err == io.EOF { // io.EOF表示文件的末尾
			break
		}
		fmt.Print(content)
	}
	



}