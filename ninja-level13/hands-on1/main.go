package main

import (
	"fmt"

	"github.com/wisetlaloc/my_go_learning/ninja-level13/hands-on1/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
