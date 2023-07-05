package test

import (
	"fmt"
)

type Person struct {
	Name string // 名字 变量名首字母大写表示公有 外界可以访问
	Age int    // 年龄
	School string // 学校
}

type Student struct {
	Person // 匿名字段
	Class string // 班级
	StudentNo int // 学号
}

func StructStu () {
	var stu Student
	stu.StudentNo = 1001
	stu.Class = "101班"
	stu.Name = "张三"
	stu.Age = 18
	stu.School = "清华大学"
	fmt.Printf("stu的值:%v\n",stu)
} 

func Struct () {
	// fmt.Println("=========结构体=========")
	fmt.Println("------------------------")
	var p1 Person
	p1.Name = "张三"
	p1.Age = 18
	p1.School = "清华大学"
	fmt.Printf("p1的值:%v\n",p1)
}

func Struct2 (Name string,Age int,School string) {
	fmt.Println("------------------------")
	var p1 Person
	p1.Name = Name
	p1.Age = Age
	p1.School = School
	fmt.Printf("p1的值:%v\n",p1)
}

func Struct3 () {
	// 创建结构体指针
	var t *Person = new(Person)
	// t是指针,指向的是一个结构体 通过指针操作结构体成员
	(*t).Name = "姗姗"
	(*t).Age = 18 // *的作用是获取指针指向的变量的值
	(*t).School = "清华大学"
	// 为了简化指针操作结构体成员,Go语言支持直接使用指针操作结构体成员
	t.School = "北京大学" // Go语言底层会自动把t转换成(*t)
	fmt.Printf("t的值:%v\n",*t)
	fmt.Println("------------------------")
}

func Struct4 () {
	// 语法糖
	var t *Person = &Person{}
	(*t).Name = "茜茜"
	t.Age = 18 // t.Age = 18 Go语言底层会自动把t转换成(*t)
	(*t).School = "清华大学"
	fmt.Printf("t的值:%v\n",*t)
}

func Struct10 () {
	// 赋值操作
	// 方式1: 按照顺序提供初始化值
	var p1 Person = Person{"张三",18,"清华大学"}
	fmt.Printf("p1的值:%v\n",p1)

	// 方式2: 通过键值对初始化,可以任意顺序
	var p2 Person = Person{
		Name:"李四",
		Age:20,
		School:"北京大学",
	}
	fmt.Printf("p2的值:%v\n",p2)

	// 方式3: 返回结构体指针,返回局部变量地址也是安全的
	var p3 *Person = &Person{"小王",18,"清华大学"}
	fmt.Printf("p3的值:%v\n",*p3)
}