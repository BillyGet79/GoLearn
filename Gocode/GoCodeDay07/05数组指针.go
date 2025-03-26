package main

import "fmt"

func main0501() {

	//定义指针 指向数组
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//自动推导类型创建数组指针
	//p := &arr

	//指针和数组建立关系
	var p *[10]int
	p = &arr

	//fmt.Printf("%T\n", p)
	//通过指针简介操作数组
	for i := 0; i < len(arr); i++ {
		fmt.Println(p[i])
	}
	//(*数组指针变量)[下标] = 值
	//(*p)[1] = 222
	//数组指针可以直接操作数组元素
	p[1] = 222 //ok
	//fmt.Println(*p) //arr
}

func main0502() {
	arr := [5]int{1, 2, 3, 4, 5}
	//指针变量和要存储的数据类型要相同

	//p1和p2在内存中指向相同的地址 但是p1和p2的类型不同
	//一个是指向整个数组 一个是指向数组的元素
	p1 := &arr
	//p2 := &arr[0]

	//fmt.Printf("%p\n", p1)
	//fmt.Printf("%p\n", p2)
	//
	//fmt.Printf("%T\n", p1)
	//fmt.Printf("%T\n", p2)

	for i := 0; i < len(p1); i++ {
		fmt.Println(p1[i])
	}
}

// 将数组指针作为函数参数
func test2(p *[5]int) {

	//通过指针简介操作数组
	p[2] = 123
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	p := &arr
	//地址传递
	test2(p)
	fmt.Println(arr)
}
