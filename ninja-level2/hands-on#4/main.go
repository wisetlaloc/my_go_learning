package main

import "fmt"

func main() {
	x := 8
	showNum(x)
	y := x << 1
	showNum(y)
}

func showNum(x int) {
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
}
