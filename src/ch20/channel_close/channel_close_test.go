package channel_close_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {

		i := 0
		for {
			i++
			ch <- i
			fmt.Println(fmt.Sprintf("生产者生产%d", i))
			time.Sleep(time.Millisecond * 100)
			if i == 20 {
				defer func() {
					fmt.Println("重复关闭channel了")
				}()

				close(ch)
			}
		}
		wg.Done()
	}()
}

func DataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		tag := 1
		for tag == 1 {
			if data, ok := <-ch; ok {
				fmt.Println(fmt.Sprintf("消费者消耗%d", data))
			} else {
				tag = 0
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	DataReceiver(ch, &wg)
	wg.Add(1)
	DataReceiver(ch, &wg)
	wg.Wait()

}
