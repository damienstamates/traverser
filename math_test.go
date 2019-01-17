package traverser

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func AddTest(t *testing.T) {
	Convey("Given two integers a and b", t, func() {
		var a, b int
		a, b = 1, 4
		Convey("When a and b are used in Add", func() {
			res := Add(a, b)
			Convey("Then then the result should equal a+b", func() {
				So(res, ShouldEqual, a+b)
			})
		})
	})
}

func SubTest(t *testing.T) {
	Convey("Given two integers a and b", t, func() {
		var a, b int
		a, b = 1, 4
		Convey("When a and b are used in Sub", func() {
			res := Add(a, b)
			Convey("Then then the result should equal a-b", func() {
				So(res, ShouldEqual, a-b)
			})
		})
	})
}
