package main

import "fmt"

func main0301() {
	//var a string = "你好"
	////fmt.Println(a)
	//fmt.Printf("%s", a)
	//a := "你好"
	//fmt.Println(a)
	//fmt.Printf("%s", a)
	//双引号引起来的称为字符串
	ch := 'a'
	str := "a" //'a''\0'字符串结束标志

	fmt.Printf("%c\n", ch)
	//%s打印字符串打印到\0之前的内容
	fmt.Printf("%s\n", str)
}

func main0302() {

	//len 函数	用来计算字符串中字符个数 不包含\0	返回值为int类型
	a := "hello world"
	count := len(a)
	fmt.Println(count)
	//在go语言中一个汉字占3个字符	为了和Linux进行统一处理
	b := "传智播客"
	count = len(b)
	fmt.Println(count)
}

func main0303() {
	str1 := "赌场上线了"
	str2 := "在线发牌"

	//字符串的连接
	str3 := str1 + str2
	fmt.Println(str3)
}

func main() {
	var str string
	_, err := fmt.Scanf("%s", &str)
	if err != nil {
		return
	}
	fmt.Printf("%s", str)
}
