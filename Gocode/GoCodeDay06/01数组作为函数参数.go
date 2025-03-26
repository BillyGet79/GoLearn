package main

import (
	"GoCodeDay06/Sort"
	"fmt"
)

func swap(a int, b int) {
	a, b = b, a
}

func main0101() {
	a := 10
	b := 20
	//值传递	形参不能改变实参的值 形参和实参是不同的地址单元
	swap(a, b)

	fmt.Println(a)
	fmt.Println(b)
}

// 数组作为函数参数
func main() {
	arr := [10]int{9, 1, 5, 6, 8, 4, 7, 10, 3, 2}
	//数组作为函数参数是值传递
	//形参和实参是不同的存储单元 内存中有两份独立的数组存储不同的数据
	//在函数调用结束 形参单元销毁 不会影响主调函数中实参的值
	//如果想通过函数计算结果并传递给实参 需要使用数组作为函数返回值
	arr = Sort.BubbleSort1(arr)

	fmt.Println(arr)
}
