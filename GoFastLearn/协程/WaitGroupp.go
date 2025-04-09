package main

import (
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 5; i++ {
		println(s)
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	go say("Hello")
	go say("World")

	wg.Wait()
}
