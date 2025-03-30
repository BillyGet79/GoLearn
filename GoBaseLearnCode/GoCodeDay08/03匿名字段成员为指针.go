package main

import "fmt"

type person1 struct {
	id   int
	name string
	age  int
	sex  string
}
type student1 struct {
	*person1 //指针作为匿名字段

	class int
	score int
}

func main() {

	var stu student1

	stu.class = 301
	stu.score = 90

	//stu.person1是一个指针类型 默认值为nil 0x0
	//需要对指针进行创建空间 new(person1)
	stu.person1 = new(person1)
	//(*stu.person1).name = "盖伦"
	stu.name = "盖伦"
	stu.id = 1
	stu.age = 30
	stu.sex = "男"

	fmt.Println(stu)

}
