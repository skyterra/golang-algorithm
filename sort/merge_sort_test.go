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

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := GenSlice(10000, 1000000)
		MergeSort(s, 0, len(s)-1)
	}
}

func BenchmarkMergeSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := GenSlice(10000, 1000000)
		MergeSort2(s, 0, len(s)-1)
	}
}
