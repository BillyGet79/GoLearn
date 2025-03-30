package main

import (
	"fmt"
	"strings"
)

func main0301() {
	//查找一个字符串在另外要给字符串中是否出现
	str1 := "hello world"
	str2 := "llo"
	//Contains(被查找字符串，查找字符串)
	//一般用作于模糊查找
	if strings.Contains(str1, str2) {
		fmt.Println("找到了")
	} else {
		fmt.Println("未找到")
	}
}

func main0302() {
	//字符串切片
	slice := []string{"1234", "5678", "2907", "2308"}
	//字符串连接
	str := strings.Join(slice, "-")
	fmt.Println(str) //1234-5678-2907-2308
}

func main0303() {
	str1 := "hello world"
	str2 := "llo" //2
	//查找一个字符串在另外一个字符串中第一次出现的位置 返回值为int类型是下标
	//没查找到返回值为-1
	i := strings.Index(str1, str2)
	fmt.Println(i)
}

func main0304() {
	str := "hello world "
	//将一个字符串重复n次
	str1 := strings.Repeat(str, 3)
	fmt.Println(str1)
}

func main0305() {
	str := "hello bad world"
	//字符串替换 用作于屏蔽敏感词汇
	//n为替换次数 如果小于0 则出现的字串全部替换
	str1 := strings.Replace(str, "bad", "good", 1)
	fmt.Println(str1)
}

func main0306() {
	str := "1234-5678-2907-2308"
	//将字符串按照标志位进行切割变成切片
	slice := strings.Split(str, "-")
	fmt.Println(slice) //[1234 5678 2907 2308]
}

func main0307() {
	str := "        a u ok    "
	//去掉字符串头尾内容 中间的会保留
	//cutset为指定的内容
	str1 := strings.Trim(str, " ")
	fmt.Printf("===%s===", str1)
}

func main0308() {
	str := "     are you ok     "
	//去掉字符串中的空格转成切片	一般用做统计单词个数
	slice := strings.Fields(str)
	fmt.Println(slice) //[are you ok]
	for _, v := range slice {
		fmt.Printf("=%s\n", v)
	}
}
