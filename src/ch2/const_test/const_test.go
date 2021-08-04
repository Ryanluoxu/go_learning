package const_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable   = 1 << iota //1
	Writable               //11
	Executable             //111
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConst2(t *testing.T) {
	a := 7 // 0111
	t.Log(a&Readable == Readable)
}
