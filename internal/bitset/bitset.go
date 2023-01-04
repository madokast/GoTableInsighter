package bitset

import (
	"fmt"
	"ti/internal/utils"
)

type word = uint64

var wordSize int = 64
var hightest word = 1 << 63
var allSet word = 0xffff_ffff_ffff_ffff

type bitset struct {
	words []word
}

func New(capacity int) bitset {
	wordNumber := capacity / wordSize
	if capacity%wordSize != 0 {
		wordNumber++
	}
	return bitset{make([]word, wordNumber)}
}

func Of(words []word) bitset {
	return bitset{words}
}

func (bs bitset) IsSubSetOf(super bitset) bool {
	min := utils.Min(len(bs.words), len(super.words))
	for wi := 0; wi < min; wi++ {
		and := bs.words[wi] & super.words[wi]
		if and != bs.words[wi] {
			return false
		}
	}
	for wi := min; wi < len(bs.words); wi++ {
		if bs.words[wi] != 0 {
			return false
		}
	}
	return true
}

func (bs bitset) NextSetBit(fromIncluding int) int {
	if len(bs.words) == 0 {
		return -1
	}
	// get the word of from
	wordIndex := fromIncluding / wordSize
	if wordIndex >= len(bs.words) {
		return -1
	}

	// get the local index in the word
	bitIndex := fromIncluding % wordSize

	var cur word
	// find next in the word
	if bs.words[wordIndex] != 0 {
		cur = bs.words[wordIndex]
		cur <<= bitIndex
		if cur != 0 {
			return fromIncluding + numberOfLeadingZeros(cur)
		}
	}

	// if not found int the word, goto next word
	wordIndex++
	for ; wordIndex < len(bs.words); wordIndex++ {
		cur = bs.words[wordIndex]
		if cur != 0 {
			return wordIndex*wordSize + numberOfLeadingZeros(cur)
		}
	}

	return -1
}

func (bs bitset) Foreach(apply func(int)) {
	for i := bs.NextSetBit(0); i > 0; i = bs.NextSetBit(i + 1) {
		apply(i)
	}
}

func (bs bitset) Highest() int {
	for i := len(bs.words) - 1; i >= 0; i-- {
		if bs.words[i] != 0 {
			return (i+1)*wordSize - numberOfTrailingZeros(bs.words[i]) - 1
		}
	}

	return -1
}

func (bs bitset) PreviousSetBit(fromIncluding int) int {
	if fromIncluding < 0 {
		return -1
	}
	if len(bs.words) == 0 {
		return -1
	}

	// word of from
	wordIndex := fromIncluding / wordSize

	// if the word is out of bound, return highest
	if wordIndex >= len(bs.words) {
		return bs.Highest()
	}

	// the local index in the word
	bitIndex := fromIncluding % wordSize

	var cur word
	// find in the word
	if bs.words[wordIndex] != 0 {
		cur = bs.words[wordIndex]
		cur >>= (wordSize - bitIndex - 1)
		if cur != 0 {
			return fromIncluding - numberOfTrailingZeros(cur)
		}
	}

	// if not found int the word, goto previous word
	wordIndex--
	for ; wordIndex >= 0; wordIndex-- {
		cur = bs.words[wordIndex]
		if cur != 0 {
			return (wordIndex+1)*wordSize - numberOfTrailingZeros(cur) - 1
		}
	}

	return -1
}

func (bs bitset) Set(i int) {
	wi := i / wordSize
	i = i % wordSize
	bs.words[wi] |= (hightest >> i)
}

func (bs bitset) Fill() {
	for i := 0; i < len(bs.words); i++ {
		bs.words[i] = allSet
	}
}

func (bs bitset) Get(i int) bool {
	wi := i / wordSize
	i = i % wordSize
	return (bs.words[wi] & (hightest >> i)) != 0
}

func (bs bitset) Copy() bitset {
	temp := make([]word, len(bs.words))
	copy(temp, bs.words)
	return bitset{temp}
}

func (bs bitset) String() string {
	nums := make([]int, 0, 8)
	for i := 0; i < len(bs.words)*wordSize; i++ {
		if bs.Get(i) {
			nums = append(nums, i)
		}
	}
	return fmt.Sprintf("bs%v", nums)
}

func numberOfLeadingZeros(i uint64) int {
	if i == 0 {
		return 64
	}
	n := 0
	for i&hightest == 0 {
		n += 1
		i <<= 1
	}
	return n
}

func numberOfTrailingZeros(i uint64) int {
	if i == 0 {
		return 64
	}
	n := 0
	for i&1 == 0 {
		n += 1
		i >>= 1
	}
	return n
}
