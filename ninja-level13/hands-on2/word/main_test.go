package word

import (
	"fmt"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(`foo bar foo foo baz baz`)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(`foo bar foo foo baz baz`)
	}
}

func ExampleCount() {
	fmt.Println(Count(`Imagine all the people, living life in peace.`))
	// Output:
	// 8
}

type countTest struct {
	data   string
	answer int
}

func TestCount(t *testing.T) {
	tests := []countTest{
		countTest{data: "And then it happened.", answer: 4},
		countTest{data: "To be, or not to be.", answer: 6},
		countTest{data: "Imagine all the people, living life in peace.", answer: 8},
	}
	for _, v := range tests {
		x := Count(v.data)
		if x != v.answer {
			t.Error("Expected ", v.answer, ", Got ", x, " with ", x)
		}
	}
}

type useCountTest struct {
	data   string
	answer map[string]int
}

func TestUseCount(t *testing.T) {
	tests := []useCountTest{
		useCountTest{data: "And then it happened.", answer: map[string]int{`And`: 1, `then`: 1, `it`: 1, `happened.`: 1}},
		useCountTest{data: "to be or not to be.", answer: map[string]int{`to`: 2, `be`: 1, `not`: 1, `or`: 1, `be.`: 1}},
		useCountTest{data: "Three blind mice. Three blind mice.", answer: map[string]int{`Three`: 2, `blind`: 2, `mice.`: 2}},
	}
	for _, val := range tests {
		x := UseCount(val.data)
		for k, v := range val.answer {
			if x[k] != v {
				t.Error("Expected: ", v, ", Got: ", x[k], ", with ", k, ", in ", val.data)
			}
		}
	}
}
