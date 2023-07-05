package test

import "fmt"

// 定义接口:接口中只能有方法的声明,不能有实现,不能有变量 定义接口的方法
type SayHi interface {
	SayHi()
}

// 接口的实现:定义结构体,实现接口中的方法
type Chinese struct {

}
// 定义Chinese的方法,具体的实现SayHi()方法
func (c Chinese) SayHi() {
	fmt.Println("你好,我是中国人")
}
func (c Chinese) niuYang() {
	fmt.Println("我会中国文化")
}


// 定义American结构体,实现SayHi()方法
type American struct {
}
func (a American) SayHi() {
	fmt.Println("Hello,I'm American")
}
func (a American) playBasketball() {
	fmt.Println("我会打篮球")
}


// 定义一个函数,参数为接口类型
func SayHiFunc(s SayHi) {
	s.SayHi()
	fmt.Println("---------断言---------")
	// ch,ok := s.(Chinese) // 类型断言
	// if ok {
	if ch,ok := s.(Chinese); ok {
		ch.niuYang()
		// fmt.Println(ch)
	} else {
		am := s.(American)
		am.playBasketball()
	}
	fmt.Println("------------------")
	switch s.(type) {
	case Chinese:
		ch := s.(Chinese)
		ch.niuYang()
	case American:
		am := s.(American)
		am.playBasketball()
	default:
		fmt.Println("我是外星人")
	}
}