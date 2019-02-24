package main

import (
	"fmt"
)

func main() {
	// Solution 1
	// c := make(chan int, 1)
	// c <- 42
	// fmt.Println(<-c)

	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}
