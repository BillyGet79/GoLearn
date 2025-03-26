package main

import "fmt"

// 结构体定义在函数外部
//
//	type 结构体名 struct {
//		//结构体成员列表
//		//成员名 数据类型
//		//name string
//		//age int
//	}

// 结构体是全局的 可以在项目所有文件使用
// 结构体是一种数据类型
type Student struct {
	id    int
	name  string
	age   int
	sex   string
	score int
	addr  string
}

func main0901() {

	//定义结构体变量 复合类型
	//var 变量名 结构体名
	//var stu Student

	//为结构体成员赋值
	//stu.name = "张大帅"
	//stu.score = 99
	//stu.arrd = "奉天皇姑"
	//stu.sex = "男"
	//stu.age = 58
	//stu.id = 1

	//定义结构体是为成员赋值
	//var stu Student = Student{1, "张宗昌", 49, "男", 5, "山东济南"}

	stu := Student{1, "孙殿英", 42, "男", 5, "北平"}
	fmt.Println(stu)
}

func main0902() {
	stu := Student{101, "朱德", 60, "男", 101, "四川"}

	fmt.Printf("%p\n", &stu)
	fmt.Printf("%p\n", &stu.id)
	//结构体成员为string 需要和结构体最大的数据类型进行对其
	fmt.Printf("%p\n", &stu.name)
	fmt.Printf("%p\n", &stu.age)
	fmt.Printf("%p\n", &stu.sex)
	fmt.Printf("%p\n", &stu.score)
	fmt.Printf("%p\n", &stu.addr)
}

func main() {
	stu := Student{102, "聂荣臻", 60, "男", 101, "湖南"}

	//将结构体变量赋值
	stu1 := stu
	//stu1.id = 103
	//
	//fmt.Println(stu1)
	//fmt.Println(stu)

	//两个结构体比较 是比较所有成员 如果成员相同 结果为真 支持== !=比较操作
	if stu1 == stu {
		fmt.Println("相同")
	} else {
		fmt.Println("不相同")
	}
}
