package map_test

import "testing"

//初始化map
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 8}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))

}

//map 的值会默认设置 int 的话默认设置为0 string 的话默认设置为空字符串
// . 判断值是否存在的话 使用 v, ok := m1[xx] 的方式
//在访问key不存在是,仍然会返回零值,不能通过返回nil 来判断元素是否存在 -- important
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	m1[3] = 4
	t.Log(m1[2])
	if v, ok := m1[3]; ok {
		t.Logf("key3 value is %d", v)
	} else {
		t.Log("key3 is not exitsting")
	}

	mString := map[int]string{}
	t.Log(mString[1])
}

// map 遍历
func TestMapTravel(t *testing.T) {
	m := map[string]string{"one": "1", "two": "2", "three": "3"}
	for k, v := range m {
		t.Log(k, v)
	}
}
