package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case a := <-cancelChan:
		fmt.Println(a, "finally Canceled")
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{} //这个相当于发送了一条消息,结合上面的方法能关闭一个cannel
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan) //广播机制 能关闭所有的channel
}

func TestCancel(t *testing.T) {

	cancelChan := make(chan struct{}, 0)

	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}

				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, cancelChan)
	}

	//要执行5次才能删除5个
	// cancel_1(cancelChan)
	// cancel_1(cancelChan)
	// cancel_1(cancelChan)
	// cancel_1(cancelChan)
	// cancel_1(cancelChan)
	cancel_2(cancelChan) //执行一次就能广播关闭所有当前的channel
	time.Sleep(time.Second * 1)
}
