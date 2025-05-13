package sort

import (
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		s := GenSlice(10, 100)
		QuickSort(s, 0, len(s)-1)

		if !IsSorted(s) {
			t.Errorf("QuickSort() = %v, want %v", s, true)
		}
	}
}

// GenSlice generates a slice of random integers.
func GenSlice(size int, max int) []int {
	s := make([]int, 0, size)

	for i := 0; i < size; i++ {
		s = append(s, rand.Intn(max))
	}

	return s
}

// IsSorted checks if the slice is sorted in ascending order.
func IsSorted(s []int) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}
