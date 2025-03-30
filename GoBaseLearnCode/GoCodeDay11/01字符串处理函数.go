package main

import (
	"fmt"
	"strings"
)

func main0101() {
	str := "helloworld"
	//是否包含该字串，返回值为bool类型
	res := strings.Contains(str, "rldw")
	if res {
		fmt.Println("找到")
	} else {
		fmt.Println("未找到")
	}
}

func main0102() {
	//字符串拼接方法
	s := []string{"1234", "2876", "3456", "3456"}
	str := strings.Join(s, "-")
	fmt.Println(str)
}

func main0103() {
	str := "helloworld"
	//子串第一个字符出现的下标，若不包含子串，则返回-1
	i := strings.Index(str, "low")
	fmt.Println(i)
}

func main0104() {
	str := "helloworld"
	//将一个字符串重复count次，count必须为合法int值，即大于等于0
	//如果为0则为空值
	s := strings.Repeat(str, 4)
	fmt.Println(s)
}

func main0105() {
	str := "helloworld"
	//将str中的old子串替换为new子串，n为替换次数，如果小于0，则为全部替换
	s := strings.Replace(str, "l", "r", -1)
	fmt.Println(s)
}

func main0106() {
	str := "1234-2876-3456-3456"
	//通过sep将str进行分割，分割成切片
	s := strings.Split(str, "-")
	fmt.Println(s)
}

func main0107() {
	str := "   hello     world   "
	//将字符串前后的cutset字符删除
	s := strings.Trim(str, " ")
	fmt.Println(s)
}

func main() {
	str := "    hel  low o rl   d"
	//去掉字符串中的空格，并返回有效数据切片
	s := strings.Fields(str)
	fmt.Println(s)
}
