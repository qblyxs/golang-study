package main

import (
    "fmt"
    "log"
    "os"

    "golang.org/x/crypto/ssh"
)

func main() {
    user := "admin"
    pass := "admin@huawei.com"
    targethost := "192.168.1.253:22"

    config := &ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{
            ssh.Password(pass),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    conn, err := ssh.Dial("tcp", targethost, config)
    if err != nil {
        log.Fatal("Failed to dial: ", err)
    }
    sess, err := conn.NewSession()
    if err != nil {
        log.Fatal("Failed to create session: ", err)
    }
    // StdinPipe方法返回一个管道，该管道将在命令启动时连接到远程命令的标准输入。
    stdin, _ := sess.StdinPipe()
    // 读取标准输入
    sess.Stdout = os.Stdout
    // 读取标准错误
    sess.Stderr = os.Stderr
    sess.Shell()
    // 通过数组定义多条命令
    cmds := []string{
        "screen-length 0 temporary",
        "display ip int brief",
        "display clock",
        "quit"}
    // 循环指向数组中的命令
    for _, cmd := range cmds {
        fmt.Fprintf(stdin, "%s\n", cmd)
    }
    // 等待命令指向完成
    sess.Wait()
    sess.Close()
}