package main

import "fmt"

// func 函数名(函数参数列表)(函数的返回值类型)
func test4(a int, b int) (sum int) {
	sum = a + b
	//fmt.Println("hello world")
	//return	表示函数的结束	如果函数有返回值return 可以将返回值返回
	return
	//fmt.Println("你愁啥")
	//fmt.Println("瞅你咋地")
}
func main0501() {
	a := 10
	b := 20
	var sum int = 0
	sum = test4(a, b)
	fmt.Println(sum)
}

// 多个返回值
func test5(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}
func main() {
	a := 10
	b := 20
	//函数有多个返回值 要一一对应接收数据
	_, sub := test5(a, b)

	fmt.Println(sub)
}
