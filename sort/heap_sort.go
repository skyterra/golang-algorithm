package sort

// HeapSort sorts the slice in ascending order using the heap sort algorithm.
func HeapSort(a []int) {
	BuildMinHeap(a)

	size := len(a)
	for size > 1 {
		a[0], a[size-1] = a[size-1], a[0]

		size--
		MinHeapify(a[:size], 0)
	}

	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
}

// Left retrieves the index of the left child of the element at index i based on 0-indexed.
func Left(i int) int {
	return 2*i + 1
}

// Right retrieves the index of the right child of the element at index i based on 0-indexed.
func Right(i int) int {
	return 2*i + 2
}

// Parent retrieves the index of the parent of the element at index i.
func Parent(i int) int {
	if i%2 == 0 {
		return (i - 2) / 2
	}

	return (i - 1) / 2
}

// BuildMinHeap builds a min-heap from the given slice.
func BuildMinHeap(a []int) {
	for i := len(a) / 2; i >= 0; i-- {
		MinHeapify(a, i)
	}
}

// MinHeapify maintains the min-heap property for the subtree rooted at index i.
func MinHeapify(a []int, i int) {
	l := Left(i)
	r := Right(i)

	minimum := i
	if l < len(a) && a[l] < a[minimum] {
		minimum = l
	}

	if r < len(a) && a[r] < a[minimum] {
		minimum = r
	}

	if minimum != i {
		a[i], a[minimum] = a[minimum], a[i]
		MinHeapify(a, minimum)
	}
}
