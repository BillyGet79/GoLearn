package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := a + b

	c += 20
	c -= 20
	c *= 20
	c /= 20
	c %= 3
	fmt.Println(c)
}
