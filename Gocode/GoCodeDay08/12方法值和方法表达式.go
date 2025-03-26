package main

import (
	"fmt"
)

type student9 struct {
	name string
	age  int
	sex  string
	addr string
}

// func()	函数类型
func (s *student9) PrintInfo() {
	fmt.Println(*s)
}

func (s *student9) EditInfo(name string, age int, sex string, addr string) {
	s.name = name
	s.sex = sex
	s.age = age
	s.addr = addr
}
func main1201() {
	stu := student9{"甄姬", 32, "女", "许昌"}
	//stu.PrintInfo()

	//函数类型定义的变量
	//对象的方法和函数都具有相同的数据类型
	f := stu.PrintInfo
	f()
	fmt.Printf("%T\n", f)
}
func main() {
	var s student9
	//通过方法为对象赋值
	s.EditInfo("武则天", 32, "女", "洛阳")
	s.PrintInfo()
}
