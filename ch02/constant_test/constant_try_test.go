package constant_test

import "testing"

const (
	Monday = 1 + iota //这个 iota 相当于是一个标识
	Tuesday
	Wednesday
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Wednesday)
}

const (
	Readable   = 1 << iota //0001   位运算
	Writable               //0010
	Executable             //0100
	Lookable
)

func TestConstantTry1(t *testing.T) {
	// a := 7 // 0111
	// a := 1
	a := 15
	t.Log(a&Readable == Readable,
		a&Writable == Writable,
		a&Executable == Executable)

	t.Log(Readable, Writable, Executable, Lookable)

}
