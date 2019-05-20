package string

import "testing"

func TestString(t *testing.T) {

	var s string
	t.Log(s) //初始化为默认值 ""
	s = "hello"
	t.Log(len(s))
	// s[1] = "3"  //string 是不可变的byte slice
	s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	// s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s))
	c := []rune(s) //能够取出这个字符串里的 Unicode
	t.Log(c)
	t.Log(len(c))
	t.Logf("中 Unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

//1 字符串使用range 的时候 使用的是rune
//2 %[1]d 表示的时候 使用后面变量的第一个值, 使用%d 展示
func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c)
	}
}
