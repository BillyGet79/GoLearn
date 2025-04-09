package main

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := (right - left) * min(height[left], height[right])
	for left < right {
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		res = max(res, (right-left)*min(height[left], height[right]))
	}
	return res
}

func main() {

}
