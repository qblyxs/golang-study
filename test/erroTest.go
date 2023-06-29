package test

import (
	"fmt"
)

func ErroTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获到异常:", err)
		}
	}()
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}