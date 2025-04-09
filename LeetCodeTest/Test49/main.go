package main

import (
	"fmt"
	"sort"
)

func sortString(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

func groupAnagrams(strs []string) [][]string {
	//key为排序好的字符串，value对应的[][]string中的下标
	m := make(map[string]int)
	var res [][]string
	for _, str := range strs {
		//首先对str进行排序处理
		temp := sortString(str)
		//看看是否存在在哈希表里
		if index, exist := m[temp]; exist {
			//如果存在，则直接在string对应下标下添加进去即可
			res[index] = append(res[index], str)
		} else {
			//如果不存在，则将其排序后的结果作为键，保存在哈希表中，并且添加到res当中
			//首先要找到现在的res的第一层的最末尾下标
			i := len(res)
			//然后添加到哈希表当中
			m[temp] = i
			res = append(res, []string{})
			res[i] = append(res[i], str)
		}
	}
	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
