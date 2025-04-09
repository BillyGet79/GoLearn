package main

import (
	"fmt"
	"time"
)

func sendCode() {
	fmt.Println("发送验证码开始")
	time.Sleep(3 * time.Second)
	fmt.Println("发送验证码完成")
}

func main() {
	//实现用户注册功能
	fmt.Println("用户注册校验完成")
	//发送验证码
	//sendCode()	//会阻塞主线程
	go sendCode()
	fmt.Println("注册验证部分完成")
	time.Sleep(4 * time.Second)
	fmt.Println("main结束")
}
