package pipefilter

import (
	"reflect"
	"testing"
)

func TestStringSpilt(t *testing.T){
	sf:= NewSpiltFilter(",")
	resp,err := sf.Process("1,2,3")
	if err != nil{
		t.Fatal(err)
	}
	//t.Log(resp)
	parts,ok := resp.([]string)
	t.Log(parts)
	t.Log(ok)
	if !ok{
		t.Fatalf("Response type is %T,but the expected is string",parts)
	}

	if !reflect.DeepEqual(parts,[]string{"1","2","3"}){
		t.Errorf("Expected value is {\"1\",\"2\",\"3\"},but actual is %v",parts)
	}
}

func TestWrongInput(t *testing.T){
	sf := NewSpiltFilter(",")
	v,err := sf.Process("haha,hehe  nihao")
	t.Log(v)
	if err!=nil{
		t.Fatal("an error is expected")
	}
}