package test

import (
	"fmt"
)

func Arry5 () {
	var arr = [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(arr)
	fmt.Printf("arr的地址:%p\n",&arr)
	fmt.Printf("arr[0]的地址:%p\n",&arr[0])
	fmt.Printf("arr[0][0]的地址:%p\n",&arr[0][0])
	fmt.Printf("arr[1]的地址:%p\n",&arr[1])
	//赋值
	fmt.Println("=========赋值=========")
	arr[1][1] = 100
	arr[0][2] = 200
	fmt.Println(arr)
	//遍历
	fmt.Println("=========遍历=========")
	//普通for循环
	for i := 0; i < len(arr); i++ {
		// arr[i]是一个数组，遍历这个数组
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("arr[%d][%d]的值为:%d\n",i,j,arr[i][j])
		}
	}
	fmt.Println("--------------------")
	//for-range循环
	for i,v := range arr {
		for j,v2 := range v {
			fmt.Printf("arr[%d][%d]的值为:%d\t",i,j,v2)
		}
		fmt.Println("\n")
	}

}