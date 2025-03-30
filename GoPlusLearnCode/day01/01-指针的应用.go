package main

import "fmt"

func test(m int) {
	var b int = 100
	b += m
	fmt.Println(b)
}

/*func main() {
	var a int = 10
	var p *int = &a
	a = 100
	fmt.Println("a = ", a)
	test(10)
	*p = 250
	fmt.Println("a = ", a)
	fmt.Println("*p = ", *p)
	a = 1000
	fmt.Println("*p = ", *p)
}*/

//func main() {
//	/*	var a int = 10
//		fmt.Println("&a", &a)*/
//	var p *int
//
//	//在heap上申请一片内存地址空间
//	p = new(int) //默认类型的默认值
//	fmt.Println(*p)
//}

func swap(a *int, b *int) {
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

func main() {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Println(a, b)
}
