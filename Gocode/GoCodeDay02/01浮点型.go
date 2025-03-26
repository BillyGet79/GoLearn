package main

import "fmt"

func main0101() {

	//浮点型数据 	分为	单精度浮点型 float32(7位)	双精度浮点型 float64(15位)
	//float64 比float32更精准
	var a float64 = 1223.456
	//保留六位小鼠数据会更精准
	fmt.Printf("%f\n", a)

	var b float32 = 3.14

	fmt.Printf("%.2f", b)

	//通过自动类型推导创建的浮点型变量 默认类型为float64
	c := 123.456

	fmt.Printf("%T\n", c)
}

//买黄瓜 3.2一斤 买三斤

func main() {
	price := 3.2
	var weight float64

	fmt.Scan(&weight)

	sum := price * weight

	//fmt.Println(sum)
	fmt.Printf("%.2f", sum)
}
