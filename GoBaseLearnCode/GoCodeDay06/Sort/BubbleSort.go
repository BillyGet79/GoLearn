package Sort

func BubbleSort1(arr [10]int) [10]int {
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

func BubbleSort2(arr []int) {
	var isFinish bool = true
	var index int = 0
	for {
		if index+1 == len(arr) {
			if isFinish {
				return
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
