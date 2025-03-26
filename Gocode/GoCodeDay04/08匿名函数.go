package main

import "fmt"

func main0801() {

	a := 10
	b := 20

	//在函数内部定义一个匿名函数
	f := func(a int, b int) {
		fmt.Println(a + b)
	}

	//fmt.Println(f)
	f(a, b)

}

func main0802() {
	a := 10
	b := 20
	//int类型
	v := func(a int, b int) int {
		return a + b
	}(a, b)
	fmt.Println(v)

	//函数类型
	w := func(a int, b int) int {
		return a + b
	}
	fmt.Println(w)
}

func main() {
	a := 10
	b := 20

	f := func() int {
		return a + b
	}

	//v := f(a, b)

	a = 100
	b = 200
	v := f()

	fmt.Println(v) //300
}
