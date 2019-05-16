package interface_test

import "testing"

type Programmer interface {
	HelloWorld() string
}

type P interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

//Action1
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

//Action2
func (g *GoProgrammer) HelloWorld() string {
	return "I am pure hello world"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer) //其实生效的是这个
	t.Log(p.HelloWorld())
}
