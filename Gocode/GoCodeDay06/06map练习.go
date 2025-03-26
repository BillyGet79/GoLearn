package main

import (
	"fmt"
)

func main() {
	//通过键盘输入一个英文字符串 统计每个字母出现的次数
	//helloworld

	var str string

	fmt.Scan(&str)
	//将字符串转成切片
	slice := []byte(str)

	//定义map进行数据存储
	//map[字符]个数
	m := make(map[byte]int)

	for i := 0; i < len(slice); i++ {
		m[slice[i]]++
	}
	fmt.Println()
	for k, v := range m {
		fmt.Printf("%c:%d ", k, v)
	}
}
