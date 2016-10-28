package tools

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTools(t *testing.T) {

	Convey("Given a list of transactions", t, func() {
		data := [][]string{
			[]string{"iphone"},
			[]string{"pencil"},
			[]string{"motorolla"},
			[]string{"bike"},
			[]string{"pencil", "beatss"},
			[]string{"iphone", "iphone-case"},
			[]string{"iphone-case", "carro", "bike", "helmet", "gloves", "jacket", "lamp"},
			[]string{"iphone", "iphone-case", "beats", "notebook", "pen", "desk", "laptop", "power bank"},
			[]string{"iphone", "iphone-case", "power bank"},
			[]string{"iphone", "iphone-case", "power bank"},
		}

		first := data[0]

		Convey("If I remove [iphone] of [iphone iphone-case]", func() {
			re := removeElement(data[0][0], data[5])
			Convey("The result array is [iphone-case]", func() {
				So(re[0], ShouldEqual, "iphone-case")
			})
		})

		Convey("When I remove duplicated element of [a  b a c]", func() {
			x := []string{"a", "b", "a", "c"}
			y := unique(x)
			Convey("The third element of the array is c", func() {
				So(y[2], ShouldEqual, "c")
			})
		})

	})
}
