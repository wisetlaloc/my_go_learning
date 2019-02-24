package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("false")
	case (2 == 2):
		fmt.Println("2==2")
	}
}
