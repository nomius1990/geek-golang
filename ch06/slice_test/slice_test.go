package slice_test

import "testing"

//切片初始化 看起开和数组差不多
func TestSliceInit(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	t.Log(s1)
}

//切片 len 和cap 之间的关系.
//结论 cap 总是 >= len .当不够的时候,如果大小是在1024 一下的话,申请2倍空间,然后1/4....
func TestSliceLenCap(t *testing.T) {
	var s1 []int
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
		t.Log(len(s1), cap(s1))
	}
}

//切片之间不能比较 切片只能和nil 比较
func TestSliceCompare(t *testing.T) {
	s1 := []int{1, 2}
	// s2 := []int{1, 2}
	// if s1 == s2 { // slice can only be compared to nil
	// 	t.Log("s1==s2")
	// }

	if s1 == nil {
		t.Log("s1 == nil")
	} else {
		t.Log(s1)
	}
}

//1切片之间共享存储 比如下面
//2 当分出来的切片超过原始切片的大小的话 则会重新分配内存,指针会变
func TestSliceMonth(t *testing.T) {
	months := []string{"First", "Second", "Third", "Fourth", "Fifth", "Sixth", "Seventh", "Eighth", "Ninth", "Tenth", "Eleventh", "Twelfth"}
	t.Log(months, len(months), cap(months))
	Q2 := months[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	Summer := months[5:8]
	t.Log(Summer, len(Summer), cap(Summer))
	Summer = append(Summer, "我是新加的")
	Summer[0] = "Unknown"
	t.Log(months, Q2, Summer)

	/*****important*********/
	Q4 := months[9:]
	Q4 = append(Q4, "13")
	Q4[0] = "Did I changed?"
	t.Log(Q4, len(Q4), cap(Q4))
	t.Log(months, len(months), cap(months))
	/*****important*********/

	// t.Log(&months[5], &Q2[2], &Summer[0]) //指针指向一个位置
	// t.Log(&months[8], &Summer[3])

	t.Log(&months[9], &months[10], &months[11])
	t.Log(&Q4[0], &Q4[1], &Q4[2])
}
