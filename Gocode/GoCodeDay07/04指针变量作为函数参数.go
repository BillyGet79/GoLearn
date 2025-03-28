package main

import "fmt"

// 指针变量作为函数参数
func swap(a *int, b *int) {
	//fmt.Println(*a)
	//fmt.Println(*b)

	//&变量 取地址操作 引用运算符
	//*指针变量 取值操作 解引用运算符
	*a, *b = *b, *a
}

func main() {

	//a := 10.9
	////通过自动推导类型创建指针变量
	////所有的指针类型都存储的是一个无符号十六进制整型数据
	//p := &a
	////*int 类型
	//fmt.Printf("%T\n", p)
	//fmt.Println(p)
	a := 10
	b := 20
	//值传递
	//swap(a, b)

	swap(&a, &b)

	fmt.Println(a, b)
}
