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
        "system-view",
        "user-interface console 0",
        "authentication-mode password", //console使用密码认证
        "set authentication password cipher admin", //设置密码为admin
        "quit",
        "user-interface vty 0 4", //设置vty 0-4为远程登录端口
        "authentication-mode aaa", //使用AAA认证
        "protocol inbound all", //允许所有协议
        "user privilege level 15", //设置用户权限为15级
        "quit",
        "ssh server-source -i Vlanif 20-22", ////假设客户端使用IP地址10.10.10.20连接服务器，该地址对应的接口为Vlanif 20。
        "stelnet server enable", //启用stelnet服务ssh服务
        "ssh user admin", //创建ssh用户admin
        "ssh user admin service-type stelnet", //设置ssh用户admin只能使用stelnet服务
        "ssh user admin authentication-type password", //设置ssh用户admin使用密码认证
        "ssh user admin password cipher admin", //设置ssh用户admin的密码为admin
        //开启http服务
        "http server enable",
        "http server-source -i MEth0/0/1"    //缺省情况下，设备默认将管理IP地址192.168.1.253配置在管理网口或VLANIF1接口下，并将该接口设置为HTTP服务器端的源接口，且设备默认未指定HTTP服务器端的IPv6源地址。
        "y",
        "quit",
        "aaa", //进入aaa视图
        "local-user admin", //创建本地用户admin
        "local-user admin password irreversible-cipher admin",    //创建本地用户admin，登录密码为admin
        "local-user admin privilege level 15",    //配置本地用户admin123的级别为15
        "y",    //确认
        "quit",
        //保存配置
        "save", //保存配置
        "y",    //确认
        "screen-length 0 temporary",
        "display ip int brief",
        "display clock",
        "quit",
    }
    // 循环指向数组中的命令
    for _, cmd := range cmds {
        fmt.Fprintf(stdin, "%s\n", cmd)
    }
    // 等待命令指向完成
    sess.Wait()
    sess.Close()
}