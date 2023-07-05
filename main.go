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

	// fmt.Println("=========数组=========")
	// test.Arry()
	// test.Arry2()
	// var arr01 = [3]int{1, 2, 3}
	// test.Arry3(arr01)
	// fmt.Printf("arr01的值:%v\n", arr01)
	// test.Arry4(&arr01)
	// fmt.Printf("arr01的值:%v\n", arr01)
	// fmt.Println("=========二维数组=========")
	// test.Arry5()
	// fmt.Println("=========切片=========")
	// test.Slice()
	// test.Slice2()
	// fmt.Println("=========map=========")
	// test.Map01()
	fmt.Println("=========结构体=========")
	test.Struct()
	fmt.Println("--------------------")
	var p2 test.Person = test.Person{}
	p2.Name = "王五"
	p2.Age = 22
	p2.School = "上海大学"
	fmt.Printf("p2的值:%v\n", p2)
	fmt.Println("--------------------")
	test.Struct2("李四", 20, "北京大学")

	fmt.Println("=========结构体指针=========")
	test.Struct3()
	test.Struct4()
	test.StructStu()
	fmt.Println("=========结构体赋值=========")
	test.Struct10()
	fmt.Println("=========方法=========")
	test.StructMethod()
	test.StructMethod2()
	fmt.Println("=========封装=========")
	p := test.NewP("丽丽")
	p.SetAge(18)
	fmt.Println(p.GetAge())
	fmt.Printf("p的值:%v\n", *p)
	fmt.Println("=========继承=========")
	//创建cat实例
	// var cat test.Cat = test.Cat{}
	cat := test.Cat{}
	cat.Age = 3
	cat.Weght = 5.6
	cat.Color = "白色"
	cat.Eat()
	cat.Move()
	cat.ShowInfo()
	cat.CatchMouse()
	fmt.Printf("cat的值:%v\n", cat)
	fmt.Println("=========接口=========")
	// 创建一个chinese实例
	var c test.Chinese = test.Chinese{}
	// 创建一个American实例
	a := test.American{}
	// 美国人打招呼
	test.SayHiFunc(a)
	// 中国人打招呼
	test.SayHiFunc(c)
	fmt.Println("=========断言=========")
	ch := test.Chinese{}
	test.SayHiFunc(ch)
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

