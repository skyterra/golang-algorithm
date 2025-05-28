package data_struct

import (
	"hash/fnv"
	"math"
	"strconv"
)

type HashFunc func(key string) uint64

type BloomFilter struct {
	bitset *Bitset
	hashs  []HashFunc
}

func (bf *BloomFilter) Add(key string) {
	for _, hashFunc := range bf.hashs {
		fingerprint := hashFunc(key)

		if !bf.bitset.Test(fingerprint) {
			bf.bitset.Set(fingerprint)
		}
	}
}

func (bf *BloomFilter) Test(key string) bool {
	for _, hashFunc := range bf.hashs {
		fingerprint := hashFunc(key)

		if !bf.bitset.Test(fingerprint) {
			return false
		}
	}

	return true
}

// NewBloomFilter create a bloom filter, n is the number of elements,
// p is the false positive rate.
func NewBloomFilter(n int, p float64) *BloomFilter {
	if p <= 0.0 || p >= 1.0 {
		panic("p must be greater than 0.0 and less than 1.0")
	}

	// bit_count = -n * log(p) / (log(2) * log(2))
	bitCount := uint(-float64(n)*math.Log(p)/(math.Log(2)*math.Log(2))) + 1

	// hash_count = bit_count * log(2) / n
	hashCount := int(float64(bitCount) * math.Log(2) / float64(n))

	hashs := make([]HashFunc, hashCount)

	for i := 0; i < hashCount; i++ {
		hashs[i] = func(key string) uint64 {
			h := fnv.New64a()
			h.Write([]byte(key + strconv.Itoa(i)))
			return h.Sum64() % uint64(bitCount)
		}
	}

	return &BloomFilter{
		bitset: NewBitset(bitCount),
		hashs:  hashs,
	}
}
