package main

import "fmt"

func trap(height []int) int {
	Lmax := height[0]
	Rmax := height[len(height)-1]
	left := 1
	right := len(height) - 2
	res := 0
	for left <= right {
		//如果左侧的最大值小于右侧的最大值，那么就先结算左边柱的雨水高度
		if Lmax < Rmax {
			//如果当前柱大于左侧最大值，那么res不做累计，更新左侧最大值
			if Lmax <= height[left] {
				Lmax = height[left]
			} else {
				res += Lmax - height[left]
			}
			left++
		} else { //此时结算右边柱的雨水高度
			if Rmax <= height[right] {
				Rmax = height[right]
			} else {
				res += Rmax - height[right]
			}
			right--
		}
	}
	return res
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
