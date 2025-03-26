package main

import "fmt"

func main0301() {

	//switch 变量 {
	//case 值1:
	//	代码体
	//case 值2:
	//	代码体
	//default:
	//	代码体
	//}

	//根据分数 >=90 A >=80 B >=70 C >=60 D 不及格 E

	var score int
	fmt.Scan(&score)
	switch score / 10 {
	case 10:
		fmt.Println("A")
	case 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}

func main0302() {

	var score int
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")

	}
}

func main0303() {
	var score int
	fmt.Scan(&score)
	switch score > 60 {

	case true:
		fmt.Println("及格")
	case false:
		fmt.Println("不及格")
	}
}

func main() {
	//根据输入的月份 计算这个月有多少天
	var year int
	var month int
	fmt.Scan(&year, &month)

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println(31)
	case 4, 6, 9, 11:
		fmt.Println(20)
	case 2:
		//判断是否是闰年 	能被4整除 但是不能被100整除 或 能被400整除
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	}
}
