package sort

// MergeSort sorts the slice in ascending order using the merge sort algorithm.
func MergeSort(a []int, p, q int) {
	if p >= q {
		return
	}

	m := p + (q-p)/2
	MergeSort(a, p, m)
	MergeSort(a, m+1, q)

	merge(a, p, m, q)
}

// merge combines two sorted subarrays a[p:m+1] and a[m+1:q+1] into a single sorted subarray a[p:q+1].
func merge(a []int, p, m, q int) {
	l, r := 0, 0

	left := make([]int, m-p+1)
	copy(left, a[p:m+1])

	right := make([]int, q-m)
	copy(right, a[m+1:q+1])

	for i := p; i < q+1; i++ {
		if l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				a[i] = left[l]
				l++
			} else {
				a[i] = right[r]
				r++
			}

			continue
		}

		if l < len(left) {
			a[i] = left[l]
			l++
		}

		if r < len(right) {
			a[i] = right[r]
			r++
		}
	}
}
