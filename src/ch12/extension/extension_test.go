package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

type Dog struct {
	Pet
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

func (d *Dog) Speak() {
	// d.p.Speak()
	fmt.Print("Wang")
}

// func (d *Dog) SpeakTo(host string) {
// 	d.Speak()
// 	// d.p.SpeakTo(host)
// 	fmt.Println(" ", host)
// }

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Chao")
	// dog.Speak()
}
