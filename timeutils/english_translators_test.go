package timeutils

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMonthToEnglish(t *testing.T) {
	Convey("Given a time.Month", t, func() {
		month := time.Month(9)
		Convey("when calling MonthToEnglish", func() {
			stringMonth := MonthToEnglish(month)
			So(stringMonth, ShouldEqual, "September")
		})
	})
}

func TestWeekdayToEnglish(t *testing.T) {
	Convey("Given a time.Weekday", t, func() {
		weekday := time.Weekday(1)
		Convey("when calling WeedayToEnglish", func() {
			stringWeekday := WeekdayToEnglish(weekday)
			So(stringWeekday, ShouldEqual, "Monday")
		})
	})
}
