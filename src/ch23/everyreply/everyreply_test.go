package everyreply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRuner := 10
	ch := make(chan string, numOfRuner) //这样不会协程泄露
	// ch := make(chan string)             //这样容易协程泄露
	for i := 0; i < numOfRuner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	return <-ch
}

func AllResponse() string {
	numOfRuner := 10
	ch := make(chan string, numOfRuner) //这样不会协程泄露
	// ch := make(chan string)             //这样容易协程泄露
	for i := 0; i < numOfRuner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	//获取到所有的消息之后返回
	finalRet := ""
	for j := 0; j < numOfRuner; j++ {
		finalRet += <-ch + "\n"
	}

	return finalRet
}

func TestFirstResponse(t *testing.T) {
	t.Log("before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("after:", runtime.NumGoroutine())
}
