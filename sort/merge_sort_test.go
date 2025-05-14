package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		s := GenSlice(10, 100)
		MergeSort(s, 0, len(s)-1)
		if !IsSorted(s) {
			t.Errorf("MergeSort() = %v, want %v", s, true)
		}
	}
}
