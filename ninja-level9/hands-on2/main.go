package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println(p.first, " ", p.last, "age", p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	saySomething(&p)
	//saySomething(p)
}
