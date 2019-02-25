// Package word provides custom functions working with strings and words
package word

import "strings"

// UseCount returns the nunmber of times each word is used in a string.
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a string.
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
