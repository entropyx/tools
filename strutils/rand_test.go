package strutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRand(t *testing.T) {
	Convey("Given a", t, func() {
		Convey("When", func() {
			s1 := Rand(10)
			s2 := Rand(10)

			Convey("udhfuda", func() {
				So(s1, ShouldHaveLength, 10)
				So(s2, ShouldHaveLength, 10)
			})

			Convey("Both strings should not be equal", func() {
				So(s1, ShouldNotEqual, s2)
			})
		})
	})
}
