package dog

import (
	"fmt"
	"testing"
)

type test struct {
	data   int
	answer int
}

func TestYears(t *testing.T) {
	tests := []test{
		test{data: 3, answer: 21},
		test{data: 4, answer: 28},
		test{data: 10, answer: 70},
	}
	for _, v := range tests {
		result := Years(v.data)
		if result != v.answer {
			t.Errorf("In value %v, expected %v, got %v", v.data, v.answer, result)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	tests := []test{
		test{data: 3, answer: 21},
		test{data: 4, answer: 28},
		test{data: 10, answer: 70},
	}
	for _, v := range tests {
		result := Years(v.data)
		if result != v.answer {
			t.Errorf("In value %v, expected %v, got %v", v.data, v.answer, result)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(4))
	// Output:
	// 28
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
