package main

import (
	"fmt"
	"os"
)

func main0601() {

	//\反斜杠在程序中表示转义字符 会将后面一个字符进行转义	[\\]表示一个[\]
	//在写路径时可以使用/来代替\
	fp, err := os.Create("a.txt")

	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	//关闭文件
	defer fp.Close()
	//将字符串写入到文件中
	//\n不会换行	原因是 在windows文本文件中换行是以\r\n操作
	//在Linux中 换行是以\n操作
	fp.WriteString("hello world\r\n")
	//不会换行
	fp.WriteString("性感荷官在线发牌")

}

func main0602() {
	fp, err := os.Create("a.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer fp.Close()

	//slice := []byte{'h', 'e', 'l', 'l', 'o'}
	//将字符写入文件中
	//fp.Write(slice)

	//slice := []byte("性感法师在线讲课")
	//
	//fp.Write(slice)
	count, err1 := fp.Write([]byte("性感法师在线讲课"))
	if err1 != nil {
		fmt.Println("写入文件失败")
		return
	} else {
		fmt.Println(count)
	}
}

func main() {
	//fp, err := os.Create("a.txt")
	//只读方式打开文件
	//fp, err := os.Open()
	//os.OpenFile(文件名,打开方式,打开权限)	如果文件不存在会报错
	//打开方式 os.O_RDONLY只读 os.O_WRONLY只写 os.O_RDWR读写 os.O_APPEND追加
	//打开权限 0-7 rwx 对应读、写、执行 （二进制推导）
	fp, err := os.OpenFile("a.txt", os.O_RDWR, 6)

	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer fp.Close()

	//获取光标流位置
	//获取从文件起始到结尾有多少个字符
	//count, _ := fp.Seek(0, os.SEEK_END)
	//count, _ := fp.Seek(0, io.SeekEnd)
	//指定位置写入
	//fp.WriteAt([]byte("hello world"), count)
	//fp.WriteAt([]byte("itcast"), 0)
	//fp.WriteAt([]byte("花儿"), 17)

	fp.WriteString("hello world\r\n")
	fp.WriteString("你愁啥\r\n")
	fp.WriteString("瞅你咋地\r\n")
	fp.WriteString("在瞅一个试试\r\n")
	fp.WriteString("对不起大哥我错了")
	//fp.WriteAt([]byte("瞅你咋地"), 13)
}
