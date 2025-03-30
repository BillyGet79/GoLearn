package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//读取文件
	fp1, err1 := os.Open("cls.jpg")
	//写入文件
	fp2, err2 := os.Create("cls1.jpg")

	if err1 != nil || err2 != nil {
		fmt.Println("操作文件失败")
		return
	}
	//关闭文件
	defer fp1.Close()
	defer fp2.Close()

	//拷贝文件
	//通过read块进行文件读取

	//通过write进行写入

	buf := make([]byte, 1024*1024)
	for {
		//将读取的字符写入到新文件中
		n, err := fp1.Read(buf)
		if err == io.EOF {
			break
		}
		fp2.Write(buf[:n])
	}
}
