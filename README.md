
## 命令
```
mkdir app
cd app
go mod init qblyxs/app

mkdir lib
cd lib
go mod init qblyxs/app/lib

cd ..
go work init ./
go work use ./lib
```


### 1. 创建目录
```
mkdir app
```
### 2. 初始化go.mod文件 go mod init
```
cd app
go mod init qblyxs/app
```
### 3. 初始化go.work文件 go work init
```
go work init ./
```
### 4. 创建引入的本地包lib
```
mkdir lib
cd lib
go mod init qblyxs/app/lib
```
#### 在lib创建hello.go文件
```
vi hello.go
```
```
package lib

import "fmt"

func SayHello() string {

	fmt.Println("hello")
	return "success"
}
```
### 5. app目录添加 lib包
```
cd ..
go work use ./lib
```
### 6. app目录创建main.go文件
```
vi main.go
```
```
package main

import (
	"fmt"
	lib "qblyxs/app/lib"
)

func main() {
	fmt.Println("This is main")
	flag := lib.SayHello()
	fmt.Println(flag)
}
```
### 6. run main.go文件
```
go run main.go 
```
```
output
This is main
hello
success
```
#### 到此，项目导入完成。
```
app
.
├── go.mod
├── go.work
├── main.go
└── lib
    ├── go.mod
    └── hello.go
```

