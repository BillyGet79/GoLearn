package main

import "fmt"

func main0801() {

	a, b, c := 0, 0, 0

	fmt.Scan(&a, &b, &c)

	sum := a + b + c
	fmt.Println(sum)
	//类型转换 数据类型(变量)	数据类型(表达式)
	//fmt.Println(float64(sum / 3))
	fmt.Printf("%f", float64(sum)/3)
}

func main() {
	var a float32 = 12.34
	var b float64 = 22.22

	//在类型转换时建议低类型转成高类型 保证数据精度
	//int8 -> int16 -> int32 -> int64
	//float32 -> float64
	//建议整型转成浮点型
	fmt.Println(float64(a) + b)

	//高类型转成低类型 可能会丢失精度 数据溢出 符号发生变化
	var c int = 1234
	fmt.Println(int8(c))
}
