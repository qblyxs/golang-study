package main

import (
	"fmt"
	"strconv"
	"time"
	"sync"
)

func main() {
	// 创建一个goroutine,去执行test()
	go test()
	// 创建一个匿名的goroutine,去执行匿名函数
	go func() {
		fmt.Println("我是一个匿名函数,我被go关键字执行了")
	}()	
	// 主函数中的代码
	for i := 0; i <= 10; i++ {
		fmt.Println("main()函数中的i:", strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
	WaitG()
}

func test() {
	for i := 0; i <= 10; i++ {
		fmt.Println("test()函数中的i:", i)
		time.Sleep(1 * time.Second)
	}
}

func WaitG () {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1) // 计数器+1
		go func (i int) {
			defer wg.Done() // 计数器-1
			fmt.Println("WaitG()函数中协程的i:", i)
		}(i)
	}
	wg.Wait() // 等待计数器变为0
}


/*
1. 什么是goroutine
goroutine是Go语言中的协程(轻量级线程实现),由Go语言的运行时(runtime)管理.在Go语言中,每一个并发的执行单元叫做一个goroutine.
2. goroutine的特点
2.1. goroutine是轻量级的线程实现,创建一个goroutine的开销很小,堆栈大小只有若干kb,可以同时创建成千上万个goroutine而不会导致系统资源衰竭.
2.2. goroutine是抢占式的调度模型,Go语言的runtime(运行时)会在合适的点进行goroutine的调度,而不是依靠Go语言的开发者在代码中显式的调用.
2.3. goroutine是基于消息的通信模型,通过channel(管道)来进行通信,而不是共享内存.
3. goroutine的使用
3.1. 创建goroutine
创建goroutine的方式有两种:
3.1.1. 使用关键字go
3.1.2. 使用go关键字加匿名函数的方式
3.2. goroutine的调度
Go语言的runtime(运行时)会在合适的点进行goroutine的调度,而不是依靠Go语言的开发者在代码中显式的调用.
3.3. goroutine的通信
goroutine之间的通信是通过channel(管道)来进行通信,而不是共享内存.
*/
