package data_struct

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	b := NewBloomFilter(100000000, 0.01)

	var numbers []uint64
	for i := 0; i < 100; i++ {
		numbers = append(numbers, 100+uint64(rand.Uint32()))
	}

	for _, n := range numbers {
		b.Add(strconv.FormatUint(n, 10))
	}

	for _, n := range numbers {
		if !b.Test(strconv.FormatUint(n, 10)) {
			t.Errorf("expected %d to be set", n)
		}
	}

	for i := 0; i < 100; i++ {
		n := uint64(rand.Uint32())
		if b.Test(strconv.FormatUint(n, 10)) {
			t.Errorf("expected %d to be unset", n)
		}
	}
}
