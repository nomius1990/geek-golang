package array_test

import "testing"

func TestArray(t *testing.T) {

	var a [3]string //先定义 后赋值
	a[0] = "Hello World"
	t.Log(a[0])
	t.Log(a[1])
	t.Log(a[2])

	c := [][]int{{1, 2}, {3, 4}} //多维数组定义
	t.Log(c)

	d := [...]int{1, 2, 3, 4, 5, 6, 7, 8} //使用 ... 来指定
	t.Log(d)
}

//遍历数组
func TestArrayTravel(t *testing.T) {

	d := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	//老式便利
	for i := 0; i < len(d); i++ {
		t.Log(d[i])
	}

	//新式遍历 只要值
	for _, e := range d {
		t.Log(e)
	}

	//包含键值
	for idx, e := range d {
		t.Log(idx, e)
	}

}

//数组截取  含头没尾巴
func TestArraySlice(t *testing.T) {
	d := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	t.Log(d[2:3])
	t.Log(d[:])
	t.Log(d[:4])
	t.Log(d[2:])
	t.Log(d[2:7])
}
