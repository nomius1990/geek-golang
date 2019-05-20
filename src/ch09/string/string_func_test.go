package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",") //字符串分隔
	t.Log(parts)
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-")) //字符串连接

}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10) //整型转换成字符串
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil { //字符串转换成整型,该函数会返回两个值
		t.Log(10 + i)
	}
}
