package main

import "fmt"

//	func 函数名(参数列表)(返回值列表){
//			代码体
//	}
//
// 函数定义	只能定义一次
// 在整个项目中函数名是唯一的 不能重名
func add(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}
func main() {

	a := 10
	b := 20

	//函数调用
	add(a, b)
	add(1, 2)
}
