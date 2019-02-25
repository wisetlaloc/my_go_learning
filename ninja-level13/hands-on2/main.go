package main

import (
	"fmt"

	"github.com/wisetlaloc/my_go_learning/ninja-level13/hands-on2/quote"
	"github.com/wisetlaloc/my_go_learning/ninja-level13/hands-on2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
