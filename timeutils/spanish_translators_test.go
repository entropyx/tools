package timeutils

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMonthToSpanish(t *testing.T) {
	Convey("Given a time.Month", t, func() {
		month := time.Month(11)
		Convey("when calling MonthToSpanish", func() {
			stringMonth := MonthToSpanish(month)
			So(stringMonth, ShouldEqual, "Noviembre")
		})
	})
}

func TestWeekdayToSpanish(t *testing.T) {
	Convey("Given a time.Weekday", t, func() {
		weekday := time.Weekday(7)
		Convey("when calling WeekdayToSpanish", func() {
			stringWeekday := WeekdayToSpanish(weekday)
			So(stringWeekday, ShouldEqual, "Domingo")
		})
	})
}
