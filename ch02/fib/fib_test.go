package fib

import "testing"
import "fmt"

// var a int

func TestFibList(t *testing.T) {

	a := 1 //var a int 1
	b := 1
	// fmt.Print(a)
	t.Log(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	// fmt.Println()
}

//go 语言可以对多个变量赋值
func TestExchange(t *testing.T) {
	a := 2
	b := 3
	a, b = b, a
	t.Log(a, b)
}
