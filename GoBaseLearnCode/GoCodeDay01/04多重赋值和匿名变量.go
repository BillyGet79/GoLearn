package main

import "fmt"

func main0401() {
	//多重赋值
	//变量个数和值的个数要一一对应
	//a, b, c := 10, 20, 30
	a, b, c, d := 10, 3.14, "传智播客", "黑马程序员"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
func main() {
	var a int = 10
	var b int = 20

	//在一个作用域范围内变量名不能重复
	//var a int = 3

	//如果在多重赋值时有新定义的变量，可以使用自动推导类型
	a, b, c, d := 110, 120, "你好", "朋友"

	fmt.Printf("a的类型是：%T\n", a)
	fmt.Printf("b的类型是：%T\n", b)
	fmt.Printf("c的类型是：%T\n", c)
	fmt.Printf("d的类型是：%T\n", d)

}
func main0403() {
	//var a int = 10
	//var b int = 20

	//_表示匿名变量 不接收数据
	_, _, c, d := 120, 110, "你好", "朋友"

	//fmt.Println(_)
	fmt.Println(c)
	fmt.Println(d)
}
