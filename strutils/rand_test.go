package strutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRand(t *testing.T) {
	Convey("When two random strings are generated", t, func() {
		s1 := Rand(10)
		s2 := Rand(10)

		Convey("They should have length 10", func() {
			So(s1, ShouldHaveLength, 10)
			So(s2, ShouldHaveLength, 10)
		})

		Convey("They should not be the same", func() {
			So(s1, ShouldNotEqual, s2)
		})
	})
}
