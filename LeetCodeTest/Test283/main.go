package main

import "fmt"

func moveZeroes(nums []int) {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] == 0 {
			fast++
		} else {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
			fast++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
