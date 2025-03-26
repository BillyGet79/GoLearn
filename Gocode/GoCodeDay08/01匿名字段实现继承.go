package main

import "fmt"

// 结构体嵌套结构体
type Person struct {
	id   int
	name string
	age  int
	sex  string
}
type Student struct {
	//将Person结构体作为student结构体的成员
	//p Person//为person结构体起名字
	Person //匿名字段实现继承关系
	class  int
	score  int
}

func main0101() {

	//创建对象
	var stu Student

	//student继承于符类person 可以使用父类的结构体成员
	stu.name = "张三"
	stu.age = 18
	stu.id = 1001
	stu.sex = "女"
	stu.class = 302
	stu.score = 99

	fmt.Println(stu)
}

func main() {
	//如果一个结构体包含其他结构体信息 需要依次进行初始化
	var stu Student = Student{Person{1001, "东方不败", 35, "不详"}, 302, 100}
	fmt.Println(stu)
}
