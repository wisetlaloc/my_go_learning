package main

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	p.first = "foo"
	p.last = "bar"
}
func main() {
	p := person{
		first: "James",
		last:  "Bond",
	}
	fmt.Println(p.first, p.last)
	changeMe(&p)
	fmt.Println(p.first, p.last)
}
