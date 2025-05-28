package data_struct

type Bitset struct {
	s []uint8
}

func (b *Bitset) Set(n uint64) {
	i, j := n/8, n%8
	b.s[i] = b.s[i] | 1<<j
}

func (b *Bitset) Unset(n uint64) {
	i, j := n/8, n%8
	b.s[i] = b.s[i] & ^(1 << j)
}

func (b *Bitset) Test(n uint64) bool {
	i := n / 8
	flag := uint8(1) << (n % 8)

	if i >= uint64(len(b.s)) {
		return false
	}

	return b.s[i]&flag > 0
}

func NewBitset(bitCount uint) *Bitset {
	b := &Bitset{
		s: make([]uint8, bitCount/8+1),
	}

	return b
}
