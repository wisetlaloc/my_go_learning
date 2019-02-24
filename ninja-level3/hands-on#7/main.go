package main

import "fmt"

func main() {
	x := "main"
	if x == "sub" {
		fmt.Println("I'm sub")
	} else if x == "main" {
		fmt.Println("I'm main")
	} else {
		fmt.Println("I'm nor main nor sub")
	}
}
