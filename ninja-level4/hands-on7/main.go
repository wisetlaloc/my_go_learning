package main

import "fmt"

func main() {
	xi1 := []string{"James", "Bond", "Shaken, not stirred"}
	xi2 := []string{"Miss", "Monneypenny", "Helloooooo, James."}
	xxi := [][]string{xi1, xi2}

	for _, v := range xxi {
		for _, w := range v {
			fmt.Println(w)
		}
	}
}
