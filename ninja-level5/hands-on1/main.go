package main

import "fmt"

func main() {
	type person struct {
		first           string
		last            string
		icecreamFavours []string
	}
	x := person{
		first:           "Bob",
		last:            "Randy",
		icecreamFavours: []string{"Vanilla", "Chocolate"},
	}
	y := person{
		first:           "Rachel",
		last:            "Adams",
		icecreamFavours: []string{"Strawberry", "Chocolate Chip"},
	}
	fmt.Println(x.first, x.last)
	for _, v := range x.icecreamFavours {
		fmt.Println(v)
	}
	fmt.Println(y.first, y.last)
	for _, v := range y.icecreamFavours {
		fmt.Println(v)
	}
}
