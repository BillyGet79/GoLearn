package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main0601() {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	var arr [10]int

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	fmt.Println(arr)

	arr1 := BubbleSort(arr[:])
	fmt.Println(arr1)

}

func main0602() {
	//猜数字 1-100
	rand.New(rand.NewSource(time.Now().UnixNano()))

	num := rand.Intn(100) + 1

	var value int
	for {
		fmt.Println("请输入一个数字：")
		fmt.Scan(&value)

		if num > value {
			fmt.Println("您输入的数字太小了！")
		} else if num < value {
			fmt.Println("您输入的数字太大了！")
		} else {
			fmt.Println("恭喜您，猜对了！")
			break
		}
	}
}

// 数组去重
func main() {
	//随机双色球彩票
	//红色 1-33	选择6个 不能重复	蓝球 1-16 选择一个 可以和红球重复

	rand.New(rand.NewSource(time.Now().UnixNano()))
	var red [6]int

	for i := 0; i < len(red); i++ {
		v := rand.Intn(33) + 1

		for j := 0; j < i; j++ {
			//数据重复
			if v == red[j] {
				//重新随机
				v = rand.Intn(33) + 1

				j = -1
			}
		}
		//将没有重复的数字添加到数组中
		red[i] = v
	}
	fmt.Println("红球：", red, "蓝球：", rand.Intn(16)+1)
}

func BubbleSort(arr []int) []int {
	var isFinish bool = true
	var index int = 0
	for {
		if index+1 == len(arr) {
			if isFinish {
				return arr
			} else {
				index = 0
				isFinish = true
				continue
			}
		}
		if arr[index] > arr[index+1] {
			isFinish = false
			temp := arr[index]
			arr[index] = arr[index+1]
			arr[index+1] = temp
		}
		index++
	}
}
