package main

import "fmt"

type human struct {
	id   int
	name string
}
type person4 struct {
	human
	age  int
	sex  string
	addr string
}
type student3 struct {
	person4
	class int
	score int
}

func main() {
	var stu student3

	stu.name = "魏璎珞"
	stu.sex = "女"
	stu.addr = "皇宫"
	stu.class = 3004
	stu.score = 100
	stu.age = 18
	stu.id = 1001

	fmt.Println(stu)
}

/*
type 技能 struct{
	名称
	范围
	CD
	消耗
	伤害
	buff
}
*/

/*
type 人物信息 struct{
	姓名
	等级
	经验
	HP
	MP
	金钱
}
*/
