package obj_pool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {

	//初始化你一个大小为10 的对象
	pool := NewObjPool(10)

	//尝试放置超出池大小的对象
	if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
		t.Error(err)
	}

	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}

	fmt.Println("Done")
}
