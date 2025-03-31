package main

import "fmt"

/*
练习：给定一个字符串列表，在原有slice上返回不包含空字符串的列表，如：
{"red", "", "black", "", "", "pink", "blue"}
--> {"red", "black", "pink", "blue"}
*/
/*func main() {
	slice := []string{"red", "", "black", "", "", "pink", "blue"}
	var res []string
	for _, s := range slice {
		if s != "" {
			res = append(res, s)
		}
	}
	fmt.Println(res)
}*/

/*
练习：写一个函数，消除[]string中重复字符串，如：
{"red", "black", "red", "pink", "blue", "pink", "blue"}
--> {"red", "black", "pink", "blue"}
*/
/*func main() {
	slice := []string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	var res []string
	set := make(map[string]struct{})
	for _, v := range slice {
		if _, exists := set[v]; !exists {
			set[v] = struct{}{}
			res = append(res, v)
		}
	}
	fmt.Println(res)
}*/

/*func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	s1 := data[8:]
	s2 := data[0:5]
	//copy(目标位置切片, 源切片)
	//拷贝过程中，直接对应位置拷贝
	copy(s2, s1)
	fmt.Println("s2 = ", s2)
	fmt.Println(data)
}*/

/*
练习：要删除slice中间的某个元素并保存现有的元素顺序。如：
{5, 6, 7, 8, 9} --> {5, 6, 8, 9}
*/
func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	data = deleteSliceElem(data, 5)
	fmt.Println(data)
}

func deleteSliceElem(data []int, i int) []int {
	slice1 := data[i:]
	slice2 := data[i+1:]
	copy(slice1, slice2)
	return data[:len(data)-1]
}
