package main

import (
	"fmt"
	"github.com/qblyxs/golang-study/lib"
	"tools"
	"test"
)



func main() {
	fmt.Println("=========匿名函数=========")
	sub := func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(1, 2))



	fmt.Println("=========异常捕获=========")
	test.ErroTest()

	fmt.Println("=========数组=========")
	// test.Arry()
	// test.Arry2()
	var arr01 = [3]int{1, 2, 3}
	test.Arry3(arr01)
	fmt.Printf("arr01的值:%v\n", arr01)
	test.Arry4(&arr01)
	fmt.Printf("arr01的值:%v\n", arr01)
	fmt.Println("=========二维数组=========")
	test.Arry5()
	fmt.Println("=========切片=========")
	test.Slice()


}

// init()函数在main()函数之前执行
// 执行顺序: 包引用 > 包变量 > 包init() > 包main() > 全局变量 > init() > main()
// 优先级：全局变量 > init() > main()

func init() {
	fmt.Println("=========初始化=========")
	fmt.Println("=========包测试=========")
	lib.SayHello()
	tools.SayTest()
	test.SayTest()
	test.SayTestB()
}

