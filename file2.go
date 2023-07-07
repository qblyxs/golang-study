package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("==========写入文件==========")
	// 写入文件
	file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件 example.txt 失败")
	}
	// 当函数退出时,关闭文件,释放资源
	defer file.Close()
	// 写入文件
	file.WriteString("第七行 使用file.WriteString进行写入\n")
	file.Write([]byte("第八行 使用file.Write进行写入操作\n"))
	fmt.Printf("文件写入成功\n")
}

