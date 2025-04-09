package main

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
func lengthOfLongestSubstring(s string) int {
	str := []byte(s)
	var bitmap [256]int
	for i := range bitmap {
		bitmap[i] = -1
	}
	N := len(str)
	ans := 1
	pre := 1
	bitmap[str[0]] = 0
	for i := 1; i < N; i++ {
		p1 := i - bitmap[str[i]]
		p2 := pre + 1
		cur := min(p1, p2)
		ans = max(ans, cur)
		pre = cur
		bitmap[str[i]] = i
	}
	return ans
}

func main() {

}
