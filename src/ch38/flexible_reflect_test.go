package flexible_reflec_test

import (
	"errors"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T){
	a := map[int]string{1:"one",2:"two",3:"three"}
	b := map[int]string{1:"one",2:"two",3:"three"}

	//t.Log(a == b) //map 只能和nil比较 ,map 和map 不能进行比较
	t.Log(reflect.DeepEqual(a,b))
	s1 := []int{1,2,3}
	s2 := []int{1,2,3}
	s3 := []int{2,3,1}
	t.Log("s1 == s2 ?",reflect.DeepEqual(s1,s2))
	t.Log("s1 == s3 ?",reflect.DeepEqual(s1,s3))
}

type Employee struct {
	EmployeeID string
	Name string `format,"normal"`
	Age int
}

func(e *Employee) UpdateAge(newVal int){
	e.Age = newVal
}

type Customer struct {
	CookieId string
	Name string
	Age int
}

func fillBySetting(st interface{},settings map[string]interface{}) error{

	//判断是不是指针
	if reflect.TypeOf(st).Kind() != reflect.Ptr{
		return errors.New("the first param should be a pointer to the struct type")
	}

	//Elem 可以获得指针所指向的结构
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct{
		return errors.New("the first param should be a pointer to the struct type")
	}

	if settings == nil{
		return errors.New("settings is nil")
	}

	var (
		field reflect.StructField
		ok bool
	)

	for k,v := range settings{
		if field,ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k);!ok{
			continue
		}

		if field.Type == reflect.TypeOf(v){
			vstr:= reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}

func TestFillNameAge(t *testing.T){
	settings := map[string]interface{}{"Name":"Mike","Age":30}
	e := Employee{}
	if err:= fillBySetting(&e,settings);err!=nil{
		t.Fatal(err)
	}

	t.Log(e)

	c:= new(Customer)
	if err:= fillBySetting(c,settings);err != nil{
		t.Fatal(err)
	}
	t.Log(*c)
}