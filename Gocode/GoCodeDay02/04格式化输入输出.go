package main

import "fmt"

func main0401() {
	fmt.Printf("35%%\n")

	//十进制整型数据
	a := 123

	//%b	占位符 打印一个数据的二进制个数
	fmt.Printf("%b\n", a)
	//%o	占位符 打印一个数据的八进制格式
	fmt.Printf("%o\n", a)

	//%x %X	占位符 打印一个数据的十六进制格式
	fmt.Printf("%x\n", a)
	fmt.Printf("%X\n", a)
	//十进制
	fmt.Printf("%d\n", a)
}

func main0402() {
	//十进制数据
	var a int = 10
	//八进制数据	八进制数据是以0开头	最大值为7
	var b int = 010
	//十六进制数据	十六进制数据是以0x开头
	var c int = 0x789

	//二进制	不能在go语言中直接表示

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func main() {
	//''引起来的只能存储一个字符
	//ch:='abc'
	ch := 'a'
	str := "瓜娃子"
	fmt.Printf("%c\n", ch)
	fmt.Printf("%s\n", str)
	fmt.Println(ch)

	//%p	占位符 打印的是一个变量对应的内存地址 是以无符号十六进制整型表示
	fmt.Printf("%p\n", &ch)
	fmt.Printf("%p\n", &str)

	a := false

	//%t	占位符 打印bool类型的值
	fmt.Printf("%t\n", a)
	fmt.Println(a)
}
