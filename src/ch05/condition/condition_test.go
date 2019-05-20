package condition_test

import "testing"

//支持在if中变量赋值
func TestIfMultiSec(t *testing.T) {
	if a := 1 == 2; a {
		t.Log("a = 2")
	} else if b := 3 == 2; b {
		t.Log("b = 2")
	} else {
		t.Log("c = 3")
	}
}

// switch 演示 -- 多个结果选项,使用逗号分隔
func TestSwitchMultiSec(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2: // 可以多个项
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("not include 0-3")
		}
	}
}

//swith 演示 -- 可以不设定 switch 之后的条件表达式,在此种情况下,整个switch 结构与多个 if...else 的逻辑作用等同
func TestSwitchCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0 && i == 2: //并且支持多个条件
			t.Log(i)
			t.Log("Even or i==3")
		case i%2 == 1:
			t.Log(i)
			t.Log("Odd")
		default:
			t.Log(i)
			t.Log("我也不知道是什么了")
		}
	}
}
