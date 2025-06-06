package main

import "fmt"

func main0701() {
	a := 10
	b := 20

	fmt.Println(a + b) //30
	fmt.Println(a - b) //-10
	fmt.Println(a * b) //200
	//两个整数相除得到的结果也是整型
	//在除法计算时 除数不能为0
	fmt.Println(a / b) //0

	//通过命令编译代码
	//go build 源文件1.go 源文件2.go
	//编译并运行程序
	//go run 源文件1.go 源文件2.go
	//异常
	/*
		1、编辑时异常
		2、编译时异常
		3、运行时异常
	*/
}

func main0702() {
	a := 10
	b := 3
	//取余运算符除数不能为0
	//取余运算符不能对浮点型使用
	c := a % b
	fmt.Println(c)
}

func main() {
	//自增自减运算符
	a := 10
	//a = a + 1
	a++ //自增 在变量本身加一
	a-- //自减

	fmt.Println(a)
	//注意：自增自减不能出现在表达式中
}
