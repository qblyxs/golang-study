package test

import (
	"fmt"
)

func Map01 () {
	fmt.Println("=========map=========")
	//map的声明
	var map1 map[string]int
	fmt.Printf("map1的值:%v\t map1的类型:%T\t map1的长度:%d\t map1的地址:%p\n",map1,map1,len(map1),&map1)

	//只声明不初始化的map内存没有分配空间，不能直接赋值
	//必须通过map的make函数进行初始化,才会分配内存空间
	//map的初始化
	map1 = make(map[string]int,10)

	map1["a"] = 1
	map1["b"] = 2
	map1["a"] = 3  //key重复，会覆盖
	fmt.Printf("map1的值:%v\n",map1)

	fmt.Println("--------------------")
	map2 := map[string]int{
		"a":1,
		"b":2,
		"c":3,}
	map2["d"] = 4
	fmt.Printf("map2的值:%v\n",map2)
	fmt.Println("=========map操作=========")
	map3 := make(map[int]string)
	fmt.Printf("map3的值:%v\n ",map3)
	//添加
	map3[1000] = "aaaa"
	map3[1001] = "bbbb"
	map3[1002] = "admin"
	fmt.Printf("添加后map3的值:%v\n ",map3)
	//修改
	map3[1001] = "admin1"
	fmt.Printf("修改后map3的值:%v\n ",map3)
	//删除
	delete(map3,1001)
	fmt.Printf("删除后map3的值:%v\n ",map3)
	//查找
	fmt.Printf("查找key为1000的值:%v\n ",map3[1000])
	//清空
	map3 = make(map[int]string)
	fmt.Printf("清空后map3的值:%v\n ",map3)
	//判断key是否存在
	//第一个返回值为key对应的value，第二个返回值为key是否存在的bool值
	value,ok := map3[1000]
	fmt.Printf("key为1000的值为:%v\n,是否存在:%v\n",value,ok)
	if ok {
		fmt.Printf("key为1000的值为:%v\n",value)
	}else {
		fmt.Println("key不存在")
	}
	//遍历 只支持for range
	for key,value := range map2 {
		fmt.Printf("key:%v\t value:%v\n",key,value)
	}

	fmt.Println("=========map嵌套=========")
	map4 := make(map[string]map[string]string)
	map4["a"] = make(map[string]string)
	map4["a"]["name"] = "张三"
	map4["a"]["age"] = "18"

	map4["b"] = make(map[string]string)
	map4["b"]["name"] = "李四"
	map4["b"]["age"] = "20"
	fmt.Printf("map4的值:%v\n",map4)
	fmt.Println("--------------------")
	for k,v := range map4 {
		fmt.Printf("key:%v\t value:%v\n",k,v)
		for k1,v1 := range v {
			fmt.Printf("key:%v\t value:%v\n",k1,v1)
		}
	}

}