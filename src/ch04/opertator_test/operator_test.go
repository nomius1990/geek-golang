package operator_test

import "testing"

const (
	Readable   = 1 << iota //0001   位运算
	Writable               //0010
	Executable             //0100
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// t.Log(a == c)
	t.Log(a == d)
}

//位运算符
// &  与
// |  或
// ^  异或
// << 往左移动若干位 (放大)
// >> 往后移动若干位 (缩小)

// &^ 按位置0运算符(只要右边的有 1 则结果为0, 如果右边的为0 那么 左边的原来是什么结果就是什么)
// 1 &^ 0 -- 1
// 1 &^ 1 -- 0
// 1 &^ 1 -- 0
// 0 &^ 0 -- 0

func TestBitClear(t *testing.T) {

	a := 7            // 0111  全是1 的话 相当于被 & 的原来是多少就是多少
	a = a &^ Readable //  0110
	a = a &^ Writable //  0110

	t.Log(a&Readable == Readable,
		a&Writable == Writable,
		a&Executable == Executable)

	// t.Log(Readable, Writable, Executable)

}
