package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type ParamFun func(op int) int

func timeSpent(inner ParamFun) ParamFun {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

// func timeSpent(inner func(op int) int) func(op int) int {
// 	return func(n int) int {
// 		start := time.Now()
// 		ret := inner(n)
// 		fmt.Println("time spent:", time.Since(start).Seconds())
// 		return ret
// 	}
// }

func TestT1(t *testing.T) {

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(1))
}

func slowFun(op int) int {
	time.Sleep(time.Duration(op) * time.Second)
	fmt.Println("op is :", op)
	fmt.Println("op is :", time.Duration(op))
	return op
}
