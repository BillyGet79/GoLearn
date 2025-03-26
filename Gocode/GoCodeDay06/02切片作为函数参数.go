package main

import (
	"GoCodeDay06/Sort"
	"fmt"
)

// 切片名作为函数参数 返回值
func test(s []int) {
	fmt.Println(s)
	fmt.Println(len(s))

	fmt.Printf("%p\n", s)
}
func main() {
	//切片名本身就是一个地址
	slice := []int{9, 1, 5, 6, 8, 4, 7, 10, 3, 2}
	fmt.Printf("%p\n", slice)
	//地址传递
	test(slice)
	Sort.BubbleSort2(slice)
	fmt.Println(slice)

}
