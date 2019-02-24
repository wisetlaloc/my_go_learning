package main

import "fmt"

func main() {
	foo()
	b1, b2 := bar()
	fmt.Println(foo(), b1, b2)
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 2, "bar called"
}
