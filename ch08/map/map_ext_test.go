package map_ext

import "testing"

//VALUE 可以是函数,实现工厂模式
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	t.Log(m[1](2), m[2](2), m[3](2))
}

//使用map 实现set
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true

	//判断元素是否存在
	// if v, ok := mySet[1]; ok {
	// 	t.Log(v)
	// }
	//判断元素是否存在
	if mySet[1] {
		t.Log(mySet[1])
	}

	//元素的个数
	t.Log(len(mySet))

	//删除元素
	delete(mySet, 1)
}
