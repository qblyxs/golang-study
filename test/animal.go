// golang继承
package test

import "fmt"

type Animal struct {
	Age int
	Weght float32
}
// 定义Animal的方法
func (a *Animal) Eat() {
	fmt.Println("动物会吃东西")
}
func (a *Animal) Move() {
	fmt.Println("动物会动")
}
func (a *Animal) ShowInfo() {
	fmt.Printf("年龄:%v 体重:%v\n",a.Age,a.Weght)
}

// 定义Cat结构体,继承Animal结构体
type Cat struct {
	Animal // 继承Animal结构体
	Color string
}
// 定义Cat的方法
func (c *Cat) CatchMouse() {
	fmt.Println("猫会抓老鼠")
}