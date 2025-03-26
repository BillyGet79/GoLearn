package main

import "fmt"

func main() {

	var arr [10]int

	//通过键盘为数组元素赋值
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(arr)

	//冒泡排序
	arr1 := BubbleSort(arr[:])
	fmt.Println(arr1)

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
