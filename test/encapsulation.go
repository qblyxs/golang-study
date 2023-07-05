package test

import "fmt"

type P struct {
	Name string
	age int // 小写表示私有 外界不能访问
}

// 通过工厂模式实现结构体的构造函数
func NewP(name string) *P {
	return &P{
		Name: name,
	}
}

// 定义set和get方法,对age字段进行封装,在函数中加入逻辑判断,保证age的合法性
func (p *P) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}
func (p *P) GetAge() int {
	return p.age
}