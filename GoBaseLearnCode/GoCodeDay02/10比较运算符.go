package main

import "fmt"

func main1001() {

	a := 10
	b := 20

	fmt.Println(a >= b) //false
	fmt.Println(a == b) //false
	fmt.Println(a != b) //true
}

func main() {
	a := 10
	b := 20

	//比较运算符返回值为bool类型
	c := a+20 > b

	fmt.Printf("%T\n", c)
	fmt.Println(c)
	//布尔类型之间不能进行比较

}
