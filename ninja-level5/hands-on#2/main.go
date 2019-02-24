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
	m := map[string]person{
		x.last: x,
		y.last: y,
	}

	for k, v := range m {
		fmt.Println(k, v.first, v.last)
		for _, fav := range v.icecreamFavours {
			fmt.Println(fav)
		}
	}
}
