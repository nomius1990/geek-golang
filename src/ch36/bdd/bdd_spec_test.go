package bdd

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("Given 2 even numbsers", t, func() {
		a := 2
		b := 4
		Convey("When add the two numbers", func() {
			c := a + b
			Convey("Then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})

		Convey("When multi the two numbers", func() {
			c := a * b
			So(c, ShouldEqual, 8)
		})
	})
}
