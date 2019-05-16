package encapsulation_test

import (
	"fmt"
	"testing"
	"unsafe"
)

//结构体 相当于定义对象
type Employee struct {
	Id   string
	Name string
	Age  int
}

// func (e Employee) String() string {
// 	fmt.Printf("no zhi zhenaddress is %x", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
// }

// 使用指针类型的的话 可以节省内存开销,少一次复制
func (e *Employee) String() string {
	fmt.Printf("address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func (e *Employee) Show() string {
	return fmt.Sprintf("这是添加的show 方法")
}

func TestInit(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e) //%T 代表类型的意思
	t.Logf("e2 is %T", e2)
}

func TestStrctOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("address is %x", unsafe.Pointer(&e.Name))
	t.Log(e.String())
	t.Log(e.Show())
}
