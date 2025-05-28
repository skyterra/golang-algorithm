package data_struct

import (
	"math"
	"math/rand"
	"testing"
)

// 使用bitmap解决大数据中判断某个数是否存在
// 题目：给 40亿个不重复的32位无符号整数， 没排过序。给一个无符号整数，如何快速判断一个数是否在这40亿个数中 ？

func TestBitset(t *testing.T) {
	b := NewBitset(math.MaxUint32)

	var numbers []uint64
	for i := 0; i < 100; i++ {
		numbers = append(numbers, uint64(rand.Uint32()))
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
