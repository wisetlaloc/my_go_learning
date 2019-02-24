package main

import "fmt"

func main() {
	fmt.Println(2 == 2 && true)
	fmt.Println(true && false)
	fmt.Println(2 == 2 || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
