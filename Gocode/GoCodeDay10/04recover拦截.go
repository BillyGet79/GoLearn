package main

import "fmt"

func demo(i int) {

	var arr [10]int

	//错误拦截到出现在错误之前
	defer func() {
		//错误拦截 panic异常错误
		err := recover()
		//判断是否出现错误
		if err != nil {
			fmt.Println(err)
		}
	}()
	//数组下标越界
	arr[i] = 123

	var p *int //0x0 nil
	//p = new(int)
	*p = 123
	fmt.Println(arr)
}

func main() {
	demo(0)
}
