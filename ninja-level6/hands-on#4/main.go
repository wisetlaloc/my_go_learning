package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p := person{
		first: "Foo",
		last:  "Bar",
		age:   24,
	}
	p.speak()
}

func (p person) speak() {
	fmt.Printf("My name is %v %v, I'm %v years old.\n", p.first, p.last, p.age)
}
