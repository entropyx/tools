package reflectutils

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testStruct struct{}

func TestDeepValue(t *testing.T) {
	Convey("Given a pointer to struct value", t, func() {
		s := &testStruct{}
		v := reflect.ValueOf(s)
		Convey("When the deep value is checked", func() {
			dv := DeepValue(v)

			Convey("The deep value should be a struct", func() {
				k := dv.Kind()
				So(k, ShouldEqual, reflect.Struct)
			})
		})
	})

	Convey("Given an interface containing a pointer struct value", t, func() {
		var iface interface{}
		s := &testStruct{}
		iface = s
		v := reflect.ValueOf(iface)
		Convey("When the deep value is checked", func() {
			dv := DeepValue(v)

			Convey("The deep value should be a struct", func() {
				k := dv.Kind()
				So(k, ShouldEqual, reflect.Struct)
			})
		})
	})
}
