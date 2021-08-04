package operator_test

import "testing"

const (
	Readable   = 1 << iota //1
	Writable               //11
	Executable             //111
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// t.Log(a == c)  // mismatched types [4]int and [5]int
	t.Log(a == d) // true
}

func TestBitClear(t *testing.T) {
	a := 7 // 0111, 可读，可写，可执行

	t.Log(a&Readable == Readable)
	// 去掉可读
	a = a &^ Readable
	t.Log(a&Readable == Readable)
	t.Log(a)
}
