package main

import "fmt"

func main() {
	arr := [5]int{3, 1, 4, 1, 5}
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", arr)
}
