package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func info(s shape) {
	switch s.(type) {
	case circle:
		fmt.Println(s.(circle).area())
	case square:
		fmt.Println(s.(square).area())
	}
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.width
}

func main() {
	s := square{
		length: 4,
		width:  2,
	}
	c := circle{
		radius: 3,
	}
	info(s)
	info(c)
}
