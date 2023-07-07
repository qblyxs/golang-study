package lib

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func NginxInit() {
	// 安装依赖
	err := runCommand("yum", "install", "-y", "git", "gcc", "gcc-c++")
	if err != nil {
		log.Fatal(err)
	}

	// 下载Nginx源码
	err = runCommand("git", "clone", "https://github.com/nginx/nginx.git")
	if err != nil {
		log.Fatal(err)
	}

	// 进入Nginx源码目录
	err = os.Chdir("nginx")
	if err != nil {
		log.Fatal(err)
	}

	// 配置编译参数
	err = runCommand("./configure", "--prefix=/usr/local/nginx")
	if err != nil {
		log.Fatal(err)
	}

	// 编译并安装Nginx
	err = runCommand("make")
	if err != nil {
		log.Fatal(err)
	}

	err = runCommand("make", "install")
	if err != nil {
		log.Fatal(err)
	}

	// 返回上级目录
	err = os.Chdir("..")
	if err != nil {
		log.Fatal(err)
	}

	// 拉取代码
	err = runCommand("git", "clone", "https://github.com/your/git-repo.git")
	if err != nil {
		log.Fatal(err)
	}

	// 进入代码目录
	err = os.Chdir("git-repo")
	if err != nil {
		log.Fatal(err)
	}

	// 安装Node.js依赖
	err = runCommand("npm", "install")
	if err != nil {
		log.Fatal(err)
	}

	// 构建代码
	err = runCommand("npm", "run", "build")
	if err != nil {
		log.Fatal(err)
	}

	// 拷贝构建结果到Nginx的网站目录
	err = runCommand("cp", "-r", "build", "/usr/local/nginx/html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("一键部署完成！")
}
