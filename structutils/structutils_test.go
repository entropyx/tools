package structutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testStruct struct {
	String      string
	Uint        uint
	Int         int
	IsSnakeCase bool
	Pointer     *string
	Struct      struct{}
}

func (t *testStruct) Func() {

}

func TestToStringMap(t *testing.T) {
	Convey("Given an struct", t, func() {
		str := "A"
		s := testStruct{
			String:      str,
			Uint:        1,
			Int:         1,
			IsSnakeCase: true,
			Pointer:     &str,
			Struct:      struct{}{},
		}
		Convey("When it is converted into a map", func() {
			m := ToStringMap(s)
			expected := map[string]string{
				"string":        "A",
				"uint":          "1",
				"int":           "1",
				"is_snake_case": "true",
			}
			Convey("The map should be a copy of the struct", func() {
				So(m, ShouldResemble, expected)
			})
		})
	})

	Convey("Given a pointer to a struct", t, func() {
		str := "A"
		s := &testStruct{
			String:      str,
			Uint:        1,
			Int:         1,
			IsSnakeCase: true,
			Pointer:     &str,
			Struct:      struct{}{},
		}
		Convey("When it is converted into a map", func() {
			m := ToStringMap(s)
			expected := map[string]string{
				"string":        "A",
				"uint":          "1",
				"int":           "1",
				"is_snake_case": "true",
			}
			Convey("The map should be a copy of the struct", func() {
				So(m, ShouldResemble, expected)
			})
		})
	})
}
