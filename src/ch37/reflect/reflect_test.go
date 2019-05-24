package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

/*
学习笔记

reflect.TypeOf  vs reflect.ValueOf

reflect.TypeOf 返回类型(reflect.Type)
reflect.ValueOf 返回值(reflect.Value)

可以从reflect.Value 获得类型
通过Kind() 来判断类型

利用反射编写灵活的代码
1 按名字访问结构的成员
reflect.ValueOf(*e).FieldByName("name)

2 按名字访问结构的方法
reflect.ValueOf(e).MethodByName("UpdateAge).Call([]reflect.Value{reflect.ValueOf(1)})

*/

func TestTypeAndValue(t *testing.T){
	var f int64 = 10
	t.Log(reflect.TypeOf(f),reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func CheckType(v interface{}){
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32,reflect.Float64:
		fmt.Println("Float")
	case reflect.Int,reflect.Int8,reflect.Int64:
		fmt.Println("Interger")
	default:
		fmt.Println("Unknown",t)
	}
}

func TestBasicType(t *testing.T){
	var f float64 = 12
	CheckType(&f)
}

type Customer struct {
	CookieId string
	Name string
	Age int
}

type Employee struct {
	EmployeeID string
	Name string `format:"normal"` //Tag
	Age int
}

func(e *Employee) UpdateAge(newVal int){
	e.Age = newVal
}

func TestDeepEqual(t *testing.T){
	a := map[int]string{1:"one",2:"two",3:"three"}
	b := map[int]string{1:"one",2:"two",3:"three"}
	t.Log("a==b?",reflect.DeepEqual(a,b))

	s1 := []int{1,2,3}
	s2 := []int{1,2,3}
	s3 := []int{2,3,1}
	t.Log("s1 == s2?",reflect.DeepEqual(s1,s2))
	t.Log("s1 == s3?",reflect.DeepEqual(s1,s3))

	c1 := Customer{"1","Mike",40}
	c2 := Customer{"1","Mike",40}
	t.Log(c1 == c2)
	t.Log(reflect.DeepEqual(c1,c2))

}

func TestInvokeByName(t *testing.T){
	e := &Employee{"1","Mike",30}

	//按照名字获取成员 --->只能拿到值
	t.Logf("Name:value(%[1]v),Type(%[1]T) ",reflect.ValueOf(*e).FieldByName("Name"))

	//拿到很多信息
	if nameField,ok := reflect.TypeOf(*e).FieldByName("Name");!ok{
		t.Error("Failed to get 'Name field.")
	}else{
		t.Log("Tag:format",nameField.Tag.Get("format"))
		t.Log(nameField) //能拿到更多的信息
		t.Log(reflect.ValueOf(*e).FieldByName("Name")) //--->只能拿到值
	}

	//按名字访问结构的方法
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Updated Age",e)
}