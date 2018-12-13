package strutils

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseCurrency(t *testing.T) {
	Convey("Given a 799.00 number", t, func() {
		s := "799.00"
		Convey("When it is converted to float", func() {
			f, _ := ParseCurrency("MXN", s)
			fmt.Println("f", f)
			Convey("The float should be 799.0", func() {
				So(f, ShouldEqual, 799.0)
			})

		})
	})

	Convey("Given a 799,00 number", t, func() {
		s := "799,00"
		Convey("When it is converted to float", func() {
			f, _ := ParseCurrency("EUR", s)
			fmt.Println("f", f)
			Convey("The float should be 799.0", func() {
				So(f, ShouldEqual, 799.0)
			})

		})
	})

	Convey("Given a 799,000.00 number", t, func() {
		s := "799,000.00"
		Convey("When it is converted to float", func() {
			f, _ := ParseCurrency("MXN", s)
			fmt.Println("f", f)
			Convey("The float should be 799000.0", func() {
				So(f, ShouldEqual, 799000.0)
			})

		})
	})

	Convey("Given a 799000.00 number", t, func() {
		s := "799000.00"
		Convey("When it is converted to float", func() {
			f, _ := ParseCurrency("MXN", s)
			fmt.Println("f", f)
			Convey("The float should be 799000.0", func() {
				So(f, ShouldEqual, 799000.0)
			})

		})
	})

	Convey("Given a 799000,00 number", t, func() {
		s := "799000,00"
		Convey("When it is converted to float", func() {
			f, _ := ParseCurrency("EUR", s)
			fmt.Println("f", f)
			Convey("The float should be 799000.0", func() {
				So(f, ShouldEqual, 799000.0)
			})

		})
	})
	Convey("Given a 799.000,00 number", t, func() {
		s := "799.000,00"
		Convey("When it is converted to float", func() {
			f, _ := ParseCurrency("EUR", s)
			fmt.Println("f", f)
			Convey("The float should be 799000.0", func() {
				So(f, ShouldEqual, 799000.0)
			})

		})
	})
}
