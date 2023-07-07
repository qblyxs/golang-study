package main

import (
	"fmt"
	"sync"
)

var (
	total int
	wg sync.WaitGroup
	lock sync.Mutex
)

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}

func add() {
	defer wg.Done()
	for i := 0; i <= 10000; i++ {
		lock.Lock() // 加锁
		total += i
		lock.Unlock() // 解锁
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i <= 10000; i++ {
		lock.Lock() // 加锁
		total -= i
		lock.Unlock() // 解锁
	}
}

