package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestT1(t *testing.T) {
	// t.Log(222)

	for i := 0; i < 10; i++ {
		go func(i int) { //这种写法很重要
			fmt.Println(i)
			time.Sleep(time.Millisecond * 1000)
		}(i)
	}
	time.Sleep(time.Millisecond * 2000)

}
