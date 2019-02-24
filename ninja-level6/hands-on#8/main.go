package main

import "fmt"

func main() {
	f := returnFunc()
	fmt.Println(f())
}

func returnFunc() func() int {
	return func() int {
		return 1
	}
}
