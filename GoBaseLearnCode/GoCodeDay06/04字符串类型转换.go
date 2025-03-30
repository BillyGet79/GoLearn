package main

import (
	"fmt"
	"strconv"
)

func main0401() {
	str := "hello world"

	//将字符串转成字符切片	强制类型转换
	slice := []byte(str)

	//fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%c", slice[i])
	}
}

func main0402() {
	//字符切片
	slice := []byte{'h', 'e', 'l', 'l', 'o'}

	//fmt.Println(slice)
	//将字符切片转换为字符串
	fmt.Println(string(slice))
}

func main0403() {
	//将其他类型转成字符串
	b := false
	str := strconv.FormatBool(b)
	fmt.Println(str)

	//base为进制数 10为十进制
	//在计算机中 进制可以表示2-36进制
	str = strconv.FormatInt(123, 10)
	fmt.Println(str)

	//f：表示转换的浮点数	prec表示保留的位数	bitSize表示32位还是64位
	str = strconv.FormatFloat(3.14159, 'f', 4, 64)
	fmt.Println(str)

	//默认转换十进制
	str = strconv.Itoa(123)
	fmt.Println(str)
}

func main0404() {
	//将字符串转成其他类型
	//err捕获错误信息
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("类型转化出错")
	} else {
		fmt.Println(b)
	}

	value, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(value)

	valuefloat, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Println(valuefloat)

	//只转换10进制
	c, _ := strconv.Atoi("123")
	fmt.Println(c)
}

func main() {

	slice := make([]byte, 0, 1024)
	//将其他类型转成字符串添加到字符切片里面

	slice = strconv.AppendBool(slice, false)
	slice = strconv.AppendInt(slice, 123, 2)
	slice = strconv.AppendFloat(slice, 3.14159, 'f', 4, 64)
	slice = strconv.AppendQuote(slice, "hello") //这里会把""保存进去
	fmt.Println(string(slice))
}
