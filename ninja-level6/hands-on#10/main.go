package main

import "fmt"

func main() {
	f := foo()
	g := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(f())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func a() {
	foo()
}
