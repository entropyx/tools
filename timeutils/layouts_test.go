package timeutils

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBigEndian(t *testing.T) {
	Convey("Given a time", t, func() {
		input := "2017-08-31"
		layout := "2006-01-02"
		tm, _ := time.Parse(layout, input)

		Convey("When it is formated to big endian", func() {
			be := BigEndian(tm, "/")

			Convey("The result should valid", func() {
				So(be, ShouldEqual, "2017/08/31")
			})
		})
	})
}

func TestLittleEndian(t *testing.T) {
	Convey("Given a time", t, func() {
		input := "2017-08-31"
		layout := "2006-01-02"
		tm, _ := time.Parse(layout, input)

		Convey("When it is formated to big endian", func() {
			be := LittleEndian(tm, "/")

			Convey("The result should valid", func() {
				So(be, ShouldEqual, "31/08/2017")
			})
		})
	})
}

func TestMiddleEndian(t *testing.T) {
	Convey("Given a time", t, func() {
		input := "2017-08-31"
		layout := "2006-01-02"
		tm, _ := time.Parse(layout, input)

		Convey("When it is formated to big endian", func() {
			be := MiddleEndian(tm, "/")

			Convey("The result should valid", func() {
				So(be, ShouldEqual, "08/31/2017")
			})
		})
	})
}
