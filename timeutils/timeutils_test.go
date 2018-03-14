package timeutils

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeUtils(t *testing.T) {
	Convey("When add month to last day of january and return february", t, func() {
		now := time.Date(2018, 1, 31, 0, 0, 0, 0, &time.Location{})
		res := AddMonth(now)
		So(res.Month(), ShouldEqual, time.February)
	})

	Convey("When add month to last day of march and return april", t, func() {
		now := time.Date(2018, 3, 31, 0, 0, 0, 0, &time.Location{})
		res := AddMonth(now)
		So(res.Month(), ShouldEqual, time.April)
	})

	Convey("When add month to last day of december and return january", t, func() {
		now := time.Date(2018, 12, 31, 0, 0, 0, 0, &time.Location{})
		res := AddMonth(now)
		So(res.Year(), ShouldEqual, 2019)
		So(res.Month(), ShouldEqual, time.January)
	})

	Convey("When add month to 12 of december and return 12 of january", t, func() {
		now := time.Date(2018, 12, 12, 0, 0, 0, 0, &time.Location{})
		res := AddMonth(now)
		So(res.Year(), ShouldEqual, 2019)
		So(res.Month(), ShouldEqual, time.January)
		So(res.Day(), ShouldEqual, 12)
	})

	Convey("When add month to 31 of november that is invalid return january", t, func() {
		now := time.Date(2018, 11, 31, 0, 0, 0, 0, &time.Location{})
		res := AddMonth(now)
		So(res.Month(), ShouldNotEqual, time.December)
	})
}
