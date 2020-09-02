package slice

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

		got := Sum(numbers)

		want := 36

		if got != want {
			t.Errorf("got %d want %d given , %v", got, want, numbers)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}
