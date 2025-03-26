package main

import "fmt"

func main() {

	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//数组逆置
	i := 0
	j := len(arr) - 1

	for i < j {

		//交换数据
		arr[i], arr[j] = arr[j], arr[i]
		//改变下标
		i++
		j--
	}

	fmt.Println(arr)

}
