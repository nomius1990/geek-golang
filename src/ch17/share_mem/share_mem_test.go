package share_mem

import (
	"sync"
	"testing"
	"time"
)

//非线性安全
func TestT1(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}

	time.Sleep(time.Second * 1)
	t.Logf("counter = %d", counter)
}

//线程安全
func TestThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}

	time.Sleep(time.Second * 1)
	t.Logf("counter = %d", counter)
}

//WaitGroup
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1) //添加等待标识
		go func() {
			defer func() {
				mut.Unlock() //开锁
			}()
			mut.Lock() //锁住
			counter++
			wg.Done() //标识结束
		}()
	}
	wg.Wait() //阻塞 等待
	t.Logf("counter = %d", counter)
}
