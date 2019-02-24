package main

import "fmt"

func main() {
	f := func() int {
		return 1
	}
	s := f()
	fmt.Println(s)
}
