package main

import "fmt"

func main() {
	xi := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	for i, v := range xi {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", xi)
}
