package main

import (
	"fmt"

	"github.com/wisetlaloc/my_go_learning/ninja-level12/hands-on1/dog"
)

type customDog struct {
	name string
	age  int
}

func main() {
	myDog := customDog{
		name: "MyDog",
		age:  dog.Years(5),
	}
	fmt.Println(myDog)
}
