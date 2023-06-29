package test

import (
	"fmt"
)

func Arry () {

	var arr01 [3]int = [3]int{1,2,3}
	fmt.Println(arr01)

	var arr02 = [3]int{1,2,3}
	fmt.Println(arr02)

	var arr03 = [...]int{1,2,3}
	fmt.Println(arr03)

	var arr04 = [...]int{1:800,0:900,2:999}
	fmt.Println(arr04)

	var arr05 = [...]string{1:"tom",0:"jack",2:"mary"}
	fmt.Println(arr05)

	fmt.Println("=========分割线=========")
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	avg := sum / len(arr)
	fmt.Println(sum,avg)
	fmt.Println("=========分割线=========")
	var arr2 [5]int16
	fmt.Println("arr2的长度为",len(arr2))
	fmt.Println("arr2的值为",arr2)
	fmt.Printf("arry2的地址为:%p\n",&arr2)
	fmt.Printf("arry2的第一个元素的地址为:%p\n",&arr2[0])
	fmt.Printf("arry2的第二个元素的地址为:%p\n",&arr2[1])

	fmt.Println("=========分割线=========")
	var scores [3]int
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请输入第%d个元素的值\n",i+1)
		fmt.Scanln(&scores[i])
	}
	fmt.Printf("该数组的值为:%d\n",scores)

	fmt.Println("=========分割线=========")
	for key,value := range scores {
		fmt.Printf("第%d个元素的值为:%d\n",key+1,value)
	}

	fmt.Println("=========分割线=========")

}