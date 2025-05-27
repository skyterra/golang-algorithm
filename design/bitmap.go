package design

// 使用bitmap解决大数据中判断某个数是否存在

// 题目：给 40亿个不重复的32位无符号整数， 没排过序。给一个无符号整数，如何快速判断一个数是否在这40亿个数中 ？

type Bitset struct {
	s []uint8
}

func (b *Bitset) Set(n uint) {
	i, j := n/8, n%8
	b.s[i] = b.s[i] | 1<<j
}

func (b *Bitset) Unset(n uint) {
	i, j := n/8, n%8
	b.s[i] = b.s[i] & ^(1 << j)
}

func (b *Bitset) Test(n uint) bool {
	i := n / 8
	flag := uint8(1) << (n % 8)

	return b.s[i]&flag > 0
}

func NewBitset(bitCount uint) *Bitset {
	b := &Bitset{
		s: make([]uint8, bitCount/8+1),
	}

	return b
}
