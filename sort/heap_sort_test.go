package sort

import (
	"math/rand"
	"testing"
)

func TestBuildMinHeap(t *testing.T) {
	a := []int{8, 6, 9, 3, 1, 5, 4, 2}
	BuildMinHeap(a)

	for i := 0; i < len(a)/2; i++ {
		if Left(i) < len(a) && a[Left(i)] < a[i] {
			t.Errorf("BuildMinHeap() left greater than parent, left:%d parent:%d", a[Left(i)], a[i])
		}

		if Right(i) < len(a) && a[Right(i)] < a[i] {
			t.Errorf("BuildMinHeap() right greater than parent, right:%d parent:%d", a[Right(i)], a[i])
		}
	}
}

func TestHeapSort(t *testing.T) {
	a := GenSlice(rand.Intn(100), rand.Intn(100))
	HeapSort(a)

	if !IsSorted(a) {
		t.Errorf("HeapSort() failed, a:%v", a)
	}
}
