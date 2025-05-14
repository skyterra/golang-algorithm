package sort

import "math/rand"

// QuickSort sorts the slice in ascending order using the quick sort algorithm.
func QuickSort(a []int, p int, q int) {
	if p >= q {
		return
	}

	k := partition(a, p, q)
	QuickSort(a, p, k-1)
	QuickSort(a, k+1, q)
}

// partition partitions the slice into two parts, where the first part is less than or equal to the pivot,
// and the second part is greater than the pivot.
// It returns the index of the pivot.
func partition(a []int, p, q int) int {
	// randomly select a pivot and move it to the first position.
	offset := rand.Intn(q - p + 1)
	a[p], a[p+offset] = a[p+offset], a[p]

	n := a[p] // pivot
	i := p    // index of the last element in the first part.

	// move the elements less than or equal to the pivot to the first part.
	for k := p + 1; k < q+1; k++ {
		if a[k] <= n {
			a[i+1], a[k] = a[k], a[i+1]
			i++
		}
	}

	// move the pivot to the correct position.
	a[p], a[i] = a[i], a[p]
	return i
}
