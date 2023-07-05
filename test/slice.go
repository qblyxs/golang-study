package test

import (
	"fmt"
)

func Slice () {
	var slice1 []int = []int{1,2,3,4,5,6,7,8,9}
	fmt.Printf("slice1的值:%v\n",slice1)
	fmt.Printf("slice1的长度:%d\n",len(slice1))
	fmt.Printf("slice1的容量:%d\n",cap(slice1))
	slice1 = append(slice1,10)
	fmt.Printf("slice1的值:%v\n",slice1)
	slice1 = slice1[0:5]
	fmt.Printf("slice1的值:%v\n",slice1)
	slice1[0] = 100
	fmt.Printf("slice1的值:%v\n",slice1)

}
