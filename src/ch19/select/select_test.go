package select_test

import (
	"fmt"
	"testing"
	"time"
)

// func TestService(t *testing.T) {
// 	fmt.Println(service())
// 	otherTask()
// }

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done")
}

func AsyncService() chan string {

	// retCh := make(chan string) //声明 一个 channel
	retCh := make(chan string, 1) //声明 一个 channel 一个消息 设定为1,这样就不会阻塞
	go func() {
		ret := service()
		fmt.Println("returned result.")

		retCh <- ret                   //往channel中放入值
		fmt.Println("service exited.") //读了值之后才会执行

	}()

	return retCh

}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) //从channel 中读值
}

//使用SELECT 实现多路选择和超时
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
		// default:
		// 	t.Log("default")
	}

}
