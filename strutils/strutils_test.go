package strutils

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStringUtils(t *testing.T) {
	Convey("Given a big number", t, func() {
		number := "10000000"
		Convey("When the number is grouped into 3 digits", func() {
			result := GroupDigits(number, ",", 3)

			Convey("The result should be 10,000,000", func() {
				So(result, ShouldEqual, "10,000,000")
			})
		})
		Convey("When the number is grouped into 2 digits", func() {
			result := GroupDigits(number, ",", 2)

			Convey("The result should be 10,00,00,00", func() {
				So(result, ShouldEqual, "10,00,00,00")
			})
		})
		Convey("When the number is grouped into 4 digits", func() {
			result := GroupDigits(number, ",", 4)

			Convey("The result should be 1000,0000", func() {
				So(result, ShouldEqual, "1000,0000")
			})
		})
		Convey("When the number is grouped into 5 digits", func() {
			result := GroupDigits(number, ",", 5)

			Convey("The result should be 100,00000", func() {
				So(result, ShouldEqual, "100,00000")
			})
		})
	})
}

func TestUrif(t *testing.T) {
	Convey("Given a uri with parameters", t, func() {
		uri := "/v1/users/:user_id/posts/:post_id/"

		Convey("When the parameters are replaced with values", func() {
			userID := 123
			postID := "abcd1234"
			newUri := Urif(uri, userID, postID)

			Convey("The new uri should include the values", func() {
				So(newUri, ShouldContainSubstring, strconv.Itoa(userID))
				So(newUri, ShouldContainSubstring, postID)
			})
		})
	})

	Convey("Given a uri with an uint32 parameter", t, func() {
		uri := "/v1/users/:user_id/"

		Convey("When the parameters are replaced with values", func() {
			userID := uint32(9194)
			newUri := Urif(uri, userID)

			Convey("The new uri should include the values", func() {
				So(newUri, ShouldEqual, "/v1/users/9194/")
			})
		})
	})
}

func TestCopy(t *testing.T) {
	Convey("Given a string", t, func() {
		str := "Hello!"

		Convey("When it is copied", func() {
			str2 := Copy(str)

			Convey("The copy should be the same", func() {
				So(str2, ShouldEqual, str)
			})
		})
	})
}

func TestToSnakeCase(t *testing.T) {
	Convey("ToSnakeCase", t, func() {
		So("i_love_golang_and_json_so_much", ShouldEqual, ToSnakeCase("ILoveGolangAndJSONSoMuch"))
		So("i_love_json", ShouldEqual, ToSnakeCase("ILoveJSON"))
		So("json", ShouldEqual, ToSnakeCase("json"))
		So("json", ShouldEqual, ToSnakeCase("JSON"))
		So("привет_мир", ShouldEqual, ToSnakeCase("ПриветМир"))
	})
}

func TestPadLeft(t *testing.T) {
	Convey("PadLeft", t, func() {
		So("01", ShouldEqual, PadLeft("1", "0", 2))
		So("0001", ShouldEqual, PadLeft("1", "0", 4))
		So("12", ShouldEqual, PadLeft("12", "0", 2))
	})
}
