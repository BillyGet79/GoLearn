package main

import "fmt"

func main() {
	/*start:
	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print("* ")
			if i == 2 && j == 1 {
				break start
			}
		}
		fmt.Println()
	}*/

	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print("* ")
			if i == 2 && j == 1 {
				goto start
			}
		}
		fmt.Println()
	}
start:
}
