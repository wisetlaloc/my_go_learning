package main

import "fmt"

const (
	a = iota + 2019
	b = iota + 2019
	c = iota + 2019
	d = iota + 2019
)

func main() {
	fmt.Println(a, b, c, d)
}
