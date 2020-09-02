package array

import (
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)

	want := 15

	if got != want {
		t.Errorf("got %d want %d given , %v", got, want, numbers)
	}
}

func BenchmarkSum(b *testing.B) {
	numbers := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}
