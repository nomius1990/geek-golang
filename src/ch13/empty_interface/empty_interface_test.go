package empty_interface

import (
	"fmt"
	"testing"
)

//空接口可以表示任何类型
//可以通过断言来讲空接口转换成指定类型

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }

	// if s, ok := p.(string); ok {
	// 	fmt.Println("string", s)
	// 	return
	// }

	// fmt.Println("Unknow Type")

	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow Type")
	}

	return
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
