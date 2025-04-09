package main

import "fmt"

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longest := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			current := 1
			for numSet[currentNum+1] {
				currentNum++
				current++
			}
			if longest < current {
				longest = current
			}
		}
	}
	return longest
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
