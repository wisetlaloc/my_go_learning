package main

import "fmt"

func main() {
	f := foo(1, 2, 3, 4, 5)
	b := bar([]int{1, 2, 3, 4, 5})
	fmt.Println(f, b)
}

func foo(vi ...int) int {
	total := 0
	for _, v := range vi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
