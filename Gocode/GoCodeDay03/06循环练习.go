package main

import (
	"fmt"
)

func main0601() {

	//敲7	7的倍数 个位为7 十位为7 需要敲桌子
	for i := 0; i <= 100; i++ {
		if i%7 == 0 || i%10 == 7 || i/10 == 7 {
			fmt.Println("敲桌子")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	//水仙花数	一个三位数 各个位数的立方和等于这个数本身
	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				if i*i*i+j*j*j+k*k*k == i*100+j*10+k {
					fmt.Println(i*100 + j*10 + k)
				}
			}
		}
	}

}
