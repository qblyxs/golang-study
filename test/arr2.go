package test

import (
	"fmt"
)

func Arry2() {
	var arr1 = [3]int{1, 2, 3}
	var arr2 = [...]int{1, 2, 3, 4}
	fmt.Printf("数组类型为:%T\n", arr1)
	fmt.Printf("数组类型为:%T\n", arr2)

}

func Arry3(arr1 [3]int) {
	arr1[0] = 100
	fmt.Printf("Arr3栈帧内arr1的值:%v\n", arr1)
}

func Arry4(arr1 *[3]int) {
	arr1[0] = 100
	fmt.Printf("Arr4栈帧内arr1的值:%v\n", arr1)
}