package test

import "fmt"

type A struct {
	Num int
}
func (a A) test() {
	fmt.Println("方法test()中 Num=",a.Num)
}
//a A 是接收者,表示test方法是属于A这个结构体的 其他类型也可以有方法
//结构体A和方法test绑定,调用方法test时,必须通过A的实例来调用
//如果其他类型变量调用test方法,编译不通过
//结构体对象传入方法时,是值传递,方法内修改结构体成员变量,不会影响原来的结构体对象

func StructMethod () {
	var a A
	a.Num = 100
	a.test()
}


type integer int
func (i integer) Print() {
	i = 100
	fmt.Println("i=",i)
}
func (i *integer) Print2() {
	*i = 100
	fmt.Println("i=",*i)
}
//方法的访问范围控制规则和函数一样
func StructMethod2 () {
	var i integer = 10
	i.Print()
	i.Print2()
	fmt.Println("i=",i)
	fmt.Println("------------------------")
}