package main

import "fmt"

func main() {
	/*
			中国古代数学家张丘在他的《算经》中提出了一个著名的”百钱百鸡问题“，
		一只公鸡值五钱，一只母鸡值三钱，三只小鸡值一钱
		现在要用百钱买百鸡，请问公鸡、母鸡、小鸡各多少只
	*/

	//cock hen chicken
	count := 0
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			count++
			checken := 100 - cock - hen
			if checken%3 == 0 && 5*cock+3*hen+checken/3 == 100 {
				fmt.Printf("公鸡：%d，母鸡：%d，小鸡：%d\n", cock, hen, checken)
			}
		}
	}
	fmt.Println("执行次数：", count)
}
