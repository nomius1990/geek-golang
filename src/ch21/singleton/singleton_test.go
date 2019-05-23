package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

func TestT1(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) //这里加入等待标识
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x \n", unsafe.Pointer(obj))
			wg.Done() //这里重置标识
		}()
	}

	wg.Wait()
}

type Singleton struct {
}

var singleInstance *Singleton //这定义一个变量,防止复用
var once sync.Once

func GetSingletonObj() *Singleton {

	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton) //第一次也是最后一次创建对象
	})

	return singleInstance
}
