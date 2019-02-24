package main

import "fmt"

func main() {
	g := func(xi []int) int {
		total := 0
		for _, v := range xi {
			total += v
		}
		return total
	}
	x := foo(g, []int{1, 2, 3, 4, 5})
	fmt.Println(x)
}

func foo(f func(xi []int) int, ii []int) int {
	n := f(ii)
	n++
	return n
}
