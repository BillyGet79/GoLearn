package main

import "fmt"

func main0901() {

	var i int = 0
	for {

		if i >= 5 {
			//跳出语句
			break
		}
		fmt.Println(i)
		i++
	}
}

func main0902() {
	sum := 0
	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			//结束本次循环，继续下次循环
			//如果在程序中儒道continue后剩余代码不会执行
			//会回到循环的位置
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println("hello world1")
	fmt.Println("hello world2")
	//如果在代码中入到goto 会跳到所定义的标志位
	//可以在一个循环中跳到另外一个循环中
	//可以在一个函数中跳到另外一个函数中
	goto FLAG
	fmt.Println("hello world3")
	fmt.Println("hello world4")
FLAG:
	fmt.Println("hello world5")
	fmt.Println("hello world6")
}
