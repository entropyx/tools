package assertions

import (
	"testing"

	"github.com/gocql/gocql"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUUIDAssertions(t *testing.T) {
	Convey("Given an uuid", t, func() {
		uuid := gocql.TimeUUID().String()

		Convey("ShouldBeUUID", func() {
			message := ShouldBeUUID(uuid)

			Convey("Should NOT return a message", func() {
				So(message, ShouldBeBlank)
			})
		})

		Convey("ShouldNotBeUUID", func() {
			message := ShouldNotBeUUID(uuid)

			Convey("Should return a message", func() {
				So(message, ShouldNotBeBlank)
			})
		})
	})

	Convey("Given any other string", t, func() {
		str := "any other string"

		Convey("ShouldBeUUID", func() {
			message := ShouldBeUUID(str)

			Convey("Should return a message", func() {
				So(message, ShouldNotBeBlank)
			})
		})

		Convey("ShouldNotBeUUID", func() {
			message := ShouldNotBeUUID(str)

			Convey("Should NOT return a message", func() {
				So(message, ShouldBeBlank)
			})
		})
	})
}
