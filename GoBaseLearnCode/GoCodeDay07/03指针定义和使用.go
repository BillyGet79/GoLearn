package main

import "fmt"

func main0301() {

	var a int = 10
	//fmt.Printf("%p\n", &a)

	//定义指针变量存储变量的地址
	var p *int = &a

	//fmt.Printf("%p\n", p)
	//通过指针间接修改变量的值
	//写操作
	*p = 123

	//fmt.Println(a)
	//读操作
	fmt.Println(*p)
}
func main() {

	//声明指针变量 默认值为0x0
	//内存地址编号为0-255的空间为系统占用 不允许用户访问(读写)
	var p *int

	fmt.Println(p)
	//new(数据类型) 开辟数据类型对应的内存空间 返回值为数据类型指针

	//gc 垃圾回收机制
	p = new(int)
	*p = 123
	fmt.Println(p)

	fmt.Println(*p)
	//fmt.Printf("%p\n", p)
	//fmt.Println(p)
	//*p = 123
	//
	//fmt.Println(*p)
}

func main0303() {

	//野指针	指针变量指向了一个未知的空间 会报错
	//var p *int = *int(0x042058088)
	//指针变量必须有一个合理的指向
	//在程序中允许出现空指针 不允许出现野指针
	//fmt.Println(*p)
}
