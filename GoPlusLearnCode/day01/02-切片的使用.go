package main

import "fmt"

/*func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	//[]内的三个属性：low,high,max
	//low：起始下标位置
	//high：结束下标位置，len=high-low
	//容量：cap=max-low
	s := arr[1:3:5]
	fmt.Println(s)
	fmt.Println(cap(s))
	fmt.Println(len(s))
}*/

/*func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := arr[2:5]
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s))

	s2 := s[2:7]
	fmt.Println("s2 = ", s2)
	fmt.Println("len(s2) = ", len(s2))
	fmt.Println("cap(s2) = ", cap(s2))
}*/

/*func main() {
	//1.自动推导赋初值
	s1 := []int{1, 2, 2, 3}
	fmt.Println("s1 = ", s1)
	//2.make创建
	//第一个参数指定类型
	//第二个参数指定len
	//第三个参数指定cap
	s2 := make([]int, 5, 10)
	fmt.Println("s2 = ", s2)
	//3.make创建不指定cap
	//此为最常用的方式
	s3 := make([]int, 10)
	fmt.Println("s3 = ", s3)
}*/

/*func main() {
	s1 := []int{1, 2, 2, 3}
	//append(切片对象， 待追加元素）
	//返回值类型为切片类型
	s1 = append(s1, 88)
	s1 = append(s1, 88)
	s1 = append(s1, 88)
	s1 = append(s1, 88)
	s1 = append(s1, 88)
	fmt.Println("s1:", s1)
}*/
