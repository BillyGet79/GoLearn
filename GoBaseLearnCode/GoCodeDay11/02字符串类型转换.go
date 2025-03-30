package main

import (
	"fmt"
	"strconv"
)

func main0201() {

	//将其他类型转成字符串
	//将bool类型转成字符串
	s := strconv.FormatBool(false)
	fmt.Println(s)

	//将int类型转换为字符串，base为指定进制
	s = strconv.FormatInt(123, 10)
	fmt.Println(s)

	//将float类型转换为字符串，prec指定显示的位数，bitSize表示处理位数
	s = strconv.FormatFloat(3.1415, 'f', -1, 64)
	fmt.Println(s)

}

func main0202() {
	str := "false"

	//将字符串转为bool类型
	s1, _ := strconv.ParseBool(str)
	fmt.Println(s1)

	//将字符串转为int类型
	str = "123"
	s2, _ := strconv.ParseInt(str, 0, 64)
	fmt.Println(s2)

	//将字符串转为float类型
	str = "3.1415"
	s3, _ := strconv.ParseFloat(str, 64)
	fmt.Println(s3)
}

func main() {
	b := make([]byte, 0, 1024)
	//将bool类型放在指定切片中
	b = strconv.AppendBool(b, false)
	//将int类型放在指定切片中
	b = strconv.AppendInt(b, 123, 10)
	//将float类型放在指定切片中
	b = strconv.AppendFloat(b, 1.234, 'f', 5, 64)
	//将字符串（包含""）放在指定切片中
	b = strconv.AppendQuote(b, "hello")
	fmt.Println(string(b))
}
