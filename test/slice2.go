package test

import (
	"fmt"
)

func Slice2 () {
	slice := make([]int8, 12, 20)

	fmt.Printf("slice的dizxhi:%d \t %p \t 长度:%d \t 容量:%d\n",slice,slice, len(slice),cap(slice))


	fmt.Printf("slice的dizxhi:%p \t %p \t 长度:%d \t 容量:%d\n",slice,&slice, len(slice),cap(slice))
	slice[0] = 55
	slice[1] = 66
	fmt.Printf("slice的值:%v\n",slice)

	slice2 := []int{1,2,3,4,5,6,7,8,9}
	slice2 = slice2[:]
	slice2 = slice2[:5]
	slice2 = slice2[2:]
	
	fmt.Printf("slice2的值:%v\t 长度:%d\t 容量:%d\n",slice2,len(slice2),cap(slice2))

	fmt.Println("-------------------")
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("slice2[%d] = %d\t",i,slice2[i])
	}
	fmt.Println("\n-------------------")
	for i, v := range slice2 {
		fmt.Printf("slice2[%d] = %d\t",i,v)
	}
	fmt.Println("\n-------------------")
	slice3 := append(slice2, 10, 11, 12, 13, 14, 15)
	fmt.Printf("slice2的地址:%p\t slice3的地址:%p\t slice3的值:%v\t 长度:%d\t 容量:%d\n",slice2,slice3,slice3,len(slice3),cap(slice3))

	
	// 切片的简写
	// slice := slice[0:5]  等价于 slice := slice[:5]
	// slice := slice[5:len(slice)] 等价于 slice := slice[5:]
	// slice := slice[0:len(slice)] 等价于 slice := slice[:]


}
