package loop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

// 无限循环
// func TestWhileAll(t *testing.T) {
// 	n := 0
// 	for {
// 		t.Log(n)
// 		n++
// 	}

// }

// 简单得循环
func TestNormalWhile(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(i)
	}
}
