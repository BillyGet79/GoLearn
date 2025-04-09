package main

func twoSum(nums []int, target int) []int {
	//key=nums值,value=对应的下标
	var m = make(map[int]int)
	for i, v := range nums {
		if index, exists := m[target-v]; exists {
			res := [2]int{i, index}
			return res[:]
		}
		m[v] = i
	}
	return nil
}

func main() {

}
