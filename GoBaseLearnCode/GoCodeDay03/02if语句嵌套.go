package main

import "fmt"

func main0201() {
	var score int

	fmt.Scan(&score)

	//嵌套循环
	if score > 700 {
		fmt.Println("我要上清华")
		if score > 720 {
			fmt.Println("我要学挖掘机")
		} else if score > 710 {
			fmt.Println("我要学美容美发")
		} else {
			fmt.Println("我要学汽修")
		}
	} else if score > 680 {
		fmt.Println("我要上北大")
		if score > 690 {
			fmt.Println("我要学盗墓")
		} else {
			fmt.Println("我要学习go语言")
		}
	}
}

// 三只小猪称体重 通过键盘输入三只小猪重量 找到最重的
func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a > b {
		//a重
		if a > c {
			fmt.Println("a最重")
		} else {
			fmt.Println("c最重")
		}
	} else {
		//b重
		if b > c {
			fmt.Println("b最重")
		} else {
			fmt.Println("c最重")
		}
	}
}
