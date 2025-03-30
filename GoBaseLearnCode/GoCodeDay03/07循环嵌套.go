package main

import (
	"fmt"
	"time"
)

func main0701() {
	//代码对齐	ctrl+alt+l
	count := 0
	//外层执行一次，内层执行一周
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			count++
			fmt.Println(i, j)
		}
	}
	fmt.Println(count)
}

// 电子时钟
func main0702() {
	//时
	for i := 0; i < 24; i++ {
		//分
		for j := 0; j < 60; j++ {
			//秒
			for k := 0; k < 60; k++ {
				time.Sleep(time.Second)
				fmt.Println(i, j, k)
			}
		}
	}

}

//func main() {
//	fmt.Println(time.Now())
//	fmt.Println(time.Now().Second())
//	fmt.Println(time.Now().Day())
//	fmt.Println(time.Now().Month())
//	fmt.Println(time.Now().Year())
//}

func main() {
	//九九乘法口诀
	/*
		1*1=1
		1*2=2	2*2=4
	*/
	//外层控制行	内层控制列
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			print(j, "*", i, "=", i*j, "\t")
		}
		println()
	}
}
