package main

import "fmt"

func main1101() {

	a := 10
	b := 20

	//c := a > b//false

	//逻辑非 ！	非真为假 非假为真
	fmt.Println(!(a > b))
	//单目运算符
	// ! ++ -- &
	//双目运算符
	// + - * / % > < == >= <=
}

func main1102() {
	a := 10
	b := 20

	//逻辑与 &&
	fmt.Println(a < b && a >= 10) //trun && true = true

}

func main1103() {
	a := 10
	b := 20
	//逻辑或	||
	fmt.Println(a < b || a > b) //true || false = true

}

func main() {
	a := 10
	b := 20

	//逻辑与的优先级高于优先或
	fmt.Println(a > b || b > a && a < 0) //false || true && false = false
}
