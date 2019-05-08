package type_test

import "testing"
import "math"

type Myint int64

//隐式类型转换测试

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64 = 2
	var c Myint
	// b = a  // 1 Go语言不允许隐式类型转换
	// c = b // 2 别名和原有类型也不能进行隐式类型转换
	t.Log(a, b, c)
}

//常用的数学常量
func TestMath(t *testing.T) {
	t.Log(math.MaxInt64)
	t.Log(math.MaxUint32)
	t.Log(math.MaxInt32)
	t.Log(math.MaxFloat64)
}

func TestPoing(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 //不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
