package mymath

import (
	"fmt"
	"testing"
)

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := []int{9000, 4, 10, 8, 6, 12}
		CenteredAvg(x)
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 1, 7, 5, 3}))
	// Output:
	// 3
}

type centeredAvgtest struct {
	data   []int
	answer float64
}

func TestCenteredAvg(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{3, 1, 5, 4, 2}
	c := []int{4, 1, 5, 4, 3, 3}
	tests := []centeredAvgtest{
		centeredAvgtest{a, 3.0},
		centeredAvgtest{b, 3.0},
		centeredAvgtest{c, 3.5},
	}
	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.answer {
			t.Error("Expected: ", v.answer, ", Got: ", x)
		}
	}
}
