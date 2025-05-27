package design

import (
	"math"
	"math/rand"
	"testing"
)

func TestBitset(t *testing.T) {
	b := NewBitset(math.MaxUint32)

	var numbers []uint
	for i := 0; i < 100; i++ {
		numbers = append(numbers, uint(rand.Uint32()))
	}

	for _, n := range numbers {
		b.Set(n)
	}

	for _, n := range numbers {
		if !b.Test(n) {
			t.Errorf("expected %d to be set", n)
		}
	}

	for _, n := range numbers {
		b.Unset(n)
	}

	for _, n := range numbers {
		if b.Test(n) {
			t.Errorf("expected %d to be unset", n)
		}
	}
}
