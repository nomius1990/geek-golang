package  sync_pool

/*
学习笔记
sync.Pool(更像是用来当做缓存用)

1 尝试从私有对象获取
2 私有对象不存在,尝试从当前processor 的共享池获取
3 如果当前processor 共享池也是空的,那么尝试去其他processor 的共享池获取
4 如果所有的子池都是空的,最后就用用户指定的New 函数产生一个新的对象返回

sync.Pool 对象的放回
1 如果私有对象不存在则保存为私有对象
2 如果私有对象存在,放入当前Processor子池的共享池中

sync.Pool 对象的生命周期
1 GC 会清除sync.pool 缓存的对象
2 对象的缓存有效期为下一次GC之前

sync.Pool 总结
1 适用于通过复用,降低复杂对象的创建和GC代价
2 协程安全,会有锁的开销
3 生命周期受GC影响,不适合做连接池等,需要自己管理生命周期的资源的池化

*/
import (
	"fmt"
	"sync"
	"testing"
)

func Test1(t *testing.T){
	t.Log(222)
}

func TestSyncPool(t *testing.T){
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 100
		},
	}

	v:=pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	//runtime.GC()  //runtime之后生命周期没了
	v1,_:=pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolInMultiGroutine(t *testing.T){

	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		wg.Add(1)
		go func(id int){
			t.Log(pool.Get())
			wg.Done()
		}(i)

	}
	wg.Wait()
}