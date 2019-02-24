package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["me"] = []string{"aa", "bb", "cc"}
	for k, v := range m {
		fmt.Println(k)
		for i, val := range v {
			fmt.Printf("\tindex: %v\tvalue: %v\n", i, val)
		}
		fmt.Println()
	}
	delete(m, "me")
	for k, v := range m {
		fmt.Println(k)
		for i, val := range v {
			fmt.Printf("\tindex: %v\tvalue: %v\n", i, val)
		}
		fmt.Println()
	}
}
