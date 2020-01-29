package errutils

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsNilFunc(t *testing.T) {
	Convey("Given a func that returns an error", t, func() {
		f := func() error {
			return errors.New("an error")
		}

		Convey("When is is passed to IsNilFunc", func() {
			ok := IsNilFunc(f)

			Convey("The result should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}
