package main

import "fmt"

// 在函数外部定义的变量 称为全局变量
// 作用域是在项目中整个文件去使用
// 定义的全局变量名不能和其他文件中的全局变量名重名
// 全局变量名可以和局部变量名重名
// 全局变量存储在内存的数据区
// 如果全局变量定义时有值 存储在初始化数据区 没有值存储在未初始化数据区
var a int = 10

func main0701() {
	//变量先定义后使用 在函数内部变量名是唯一的
	//在函数内部定义的变量 称为局部变量
	//局部变量的作用域在函数的内部
	//a := 10
	//fmt.Println(a)
	var i int = 10
	for i = 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)
	//在go语言中会采用就近原则 如果在函数内部定义局部变量 和全局变量名重名 会使用局部变量
	a := 123

	//修改全局变量的值
	//如果全局变量的值修改 会影响其他位置使用全局变量的值
	a = 110
	fmt.Println(a)
	demo3()
}

func demo3() {
	fmt.Println(a)
}

func main() {

	//打印代码区的地址	代码区
	fmt.Println(demo3)
	//打印全局变量的地址	数据区
	fmt.Println(&a)
	a := 10
	//打印局部变量的地址	栈区
	fmt.Println(&a)
}
