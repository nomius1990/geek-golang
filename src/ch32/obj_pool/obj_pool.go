package obj_pool

import (
	"errors"
	"time"
)

type ReusableObj struct {
}
type ObjPool struct {
	bufChan chan *ReusableObj
}

//创建指定大小的对象池
func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}

	return &objPool
}

//给对象添加方法 参考 CH12
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): //超时控制
		return nil, errors.New("time out")
	}

}

//给对象添加方法 参考 CH12
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
