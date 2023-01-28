package bitset

import "testing"

func TestBitset01(t *testing.T) {
	bs := New(64)
	bs.Set(5)
	bs.Set(15)
	bs.Set(25)
	t.Log(bs)
}

func TestBitset02(t *testing.T) {
	bs := New(512)
	bs.Set(50)
	bs.Set(150)
	bs.Set(250)
	bs.Set(511)
	t.Log(bs)
}

func TestBitset03(t *testing.T) {
	bs := New(64)
	bs.Fill()
	t.Log(bs)
}

func TestBitset_IsSubSetOf(t *testing.T) {
	bs1 := New(64)
	bs2 := New(64)

	bs1.Set(0)
	bs1.Set(4)

	bs2.Set(0)
	bs2.Set(4)

	if !bs1.IsSubSetOf(bs2) {
		t.Error("IsSubSetOf Error")
	}

	if !bs2.IsSubSetOf(bs1) {
		t.Error("IsSubSetOf Error")
	}

	bs2.Set(10)

	if !bs1.IsSubSetOf(bs2) {
		t.Error("IsSubSetOf Error")
	}

	if bs2.IsSubSetOf(bs1) {
		t.Error("IsSubSetOf Error")
	}
}

func TestBitset04(t *testing.T) {
	bs := New(512)
	bs.Set(50)
	bs.Set(150)
	bs.Set(250)
	bs.Set(511)
	t.Log(bs)

	for i := bs.NextSetBit(0); i > 0; i = bs.NextSetBit(i + 1) {
		t.Log("NextSetBit:", i)
	}

	for i := bs.PreviousSetBit(512); i > 0; i = bs.PreviousSetBit(i - 1) {
		t.Log("PreviousSetBit:", i)
	}

	bs.Foreach(func(i int) { t.Log("Foreach", i) })
}

func TestBitsetCopy(t *testing.T) {
	bs := New(64)
	bs.Set(1)
	bs.Set(2)
	bs.Set(3)
	bs2 := bs.Copy()
	bs2.Set(4)
	bs2.Set(5)
	bs2.Set(6)
	t.Log("Origin", bs)
	t.Log("Copy", bs2)
}

func TestBitMove(t *testing.T) {
	var a uint64 = 0
	a--
	t.Log("-1 for unint64", a)
	a >>= 1
	t.Log("-1 >>1 for unint64", a)
}

func TestNumberOfLeadingZeros(t *testing.T) {
	var a uint64 = 1
	n := numberOfLeadingZeros(a)
	if n != 63 {
		t.Error("n != 63", n)
	}
	for n > 0 {
		n--
		a <<= 1
		if n != numberOfLeadingZeros(a) {
			t.Error(numberOfLeadingZeros(a), "!=", n)
		}
		t.Log(a, "numberOfLeadingZeros", numberOfLeadingZeros(a))
	}
}

func TestNumberOfTrailingZeros(t *testing.T) {
	var a uint64 = 1
	n := numberOfTrailingZeros(a)
	if n != 0 {
		t.Error("n != 0", n)
	}
	for n < 64 {
		n++
		a <<= 1
		if n != numberOfTrailingZeros(a) {
			t.Error(numberOfTrailingZeros(a), "!=", n)
		}
		t.Log(a, "numberOfTrailingZeros", numberOfTrailingZeros(a))
	}
}
