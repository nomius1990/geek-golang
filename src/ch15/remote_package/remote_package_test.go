package remote_package

import (
	cm "github.com/easierway/concurrent_map" //起别名
	"testing"
)

func TestT1(t *testing.T) {
	// t.Log("111")
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
