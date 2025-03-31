package main

import (
	"fmt"
	"strings"
)

/*func main() {

	//1.var定义
	var m1 map[int]string
	fmt.Println(m1 == nil)
	//此定义方法定义后并不能存储键值对，因为其变量是一个nil，即为空
	//m1[100] = "tiny" //err

	//2.赋值定义
	m2 := map[int]string{}
	fmt.Println(len(m2))
	//此时就可以插入值了
	m2[4] = "red"
	fmt.Println("m2 = ", m2)
	//3.make定义
	m3 := make(map[int]string)
	fmt.Println(len(m3))
	fmt.Println("m3 = ", m3)
	//4.make定义指定大小
	m4 := make(map[int]string, 10)
	fmt.Println(len(m4))
	fmt.Println("m4 = ", m4)

	//初始化map
	//1.定义同时初始化
	var m5 map[int]string = map[int]string{1: "Luffy", 250: "Sanji", 130: "Zero"}
	fmt.Println("m5 = ", m5)
	//2.自动推导类型
	m6 := map[int]string{1: "Luffy", 250: "Sanji", 130: "Zero"}
	fmt.Println("m6 = ", m6)

}*/

/*func main() {
	m1 := make(map[int]string, 1)
	m1[100] = "Nami"
	m1[20] = "Hello"
	m1[3] = "World"
	fmt.Println("m1:", m1)

	m1[3] = "yellow" //成功，将原map中key值为3的map元素替换
	fmt.Println("m1:", m1)
}*/

/*func main() {

	m1 := make(map[int]string, 1)
	m1[100] = "Nami"
	m1[20] = "Hello"
	m1[3] = "World"
	//遍历map
	for k, v := range m1 {
		fmt.Println("key:", k, "value:", v)
	}

	//判断map中的key是否存在
	if value, exists := m1[3]; !exists {
		fmt.Println("value=", value, "exists=", exists)
	} else {
		fmt.Println("value=", value, "exists=", exists)
	}
}*/

/*func main() {
	m1 := make(map[int]string, 1)
	m1[100] = "Nami"
	m1[20] = "Hello"
	m1[3] = "World"

	//map做函数参数、返回值，传引用
	m1 = mapDelete(m1, 20)
	fmt.Println(m1)
}

func mapDelete(m map[int]string, key int) map[int]string {
	//该方法没有任何返回值，直接删除使用即可
	delete(m, key)
	return m
}*/

/*
练习：封装wcFunc()函数。接收一段英文字符串str。返回一个map，记录str中每个"词"出现的次数
如："I love you and I love my family too"
*/
func main() {
	str := "I love my work and I love my family too"
	m := wcFunc(str)
	fmt.Println(m)
}

func wcFunc(str string) map[string]int {
	slice := strings.Split(str, " ")
	res := make(map[string]int)
	for _, s := range slice {
		if value, exists := res[s]; exists {
			res[s] = value + 1
		} else {
			res[s] = 1
		}
	}
	return res
}
