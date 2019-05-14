package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestT1(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(1))
}

func slowFun(op int) int {

	time.Sleep(time.Duration(op) * time.Second)
	fmt.Println("op is :", op)
	fmt.Println("op is :", time.Duration(op))
	return op
}

//返回多个值
func returnMultiValues() (int, int) {
	// return 2, 3
	return rand.Intn(10), rand.Intn(30)
}

//传入的参数可以是函数 ,返回的参数也可以是函数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

//函数可以作为变量的值
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestChangeParams(t *testing.T) {

	t.Log(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

func Clear() {
	fmt.Println("world")
}

//可变参数
func sum(ops ...int) int {
	s := 0
	for _, op := range ops {
		s += op
	}
	return s
}

func TestDefer(t *testing.T) {

	defer func() {
		// t.Log("world")
		fmt.Println("world")
	}()

	// defer Clear()

	// t.Log("hello")
	fmt.Println("hello")
	panic("Fatal error") //即使抛出致命错误 仍然会执行

}
