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
			[]string{"iphone", "pencil", "beatss", "power bank"},
		}

		// first := data[0]

		Convey("If I remove [iphone] of [iphone iphone-case]", func() {
			re := RemoveElement(data[0][0], data[5])
			Convey("The result array is [iphone-case]", func() {
				So(re[0], ShouldEqual, "iphone-case")
			})
		})

		Convey("If I remove the 0 index of data", func() {
			re := RemoveArray(1, data)
			Convey("The new 0 array should be [pencil]", func() {
				So(re[1], ShouldResemble, []string{"motorolla"})
			})
		})

		Convey("For validate that pencil is subset of [pencil beats]", func() {
			ss := Subset(data[4], data[10])
			Convey("The output should be true", func() {
				So(ss, ShouldEqual, true)
			})
		})

		Convey("When I remove duplicated element of [a  b a c]", func() {
			x := []string{"a", "b", "a", "c"}
			y := Unique(x)
			Convey("The third element of the array is c", func() {
				So(y[2], ShouldEqual, "c")
			})
		})

	})

	Convey("Given two different slices", t, func() {
		slice1 := []string{"1", "2", "3"}
		slice2 := []string{"1", "2", "3", "4"}
		Convey("It should return the difference between the two", func() {
			result := []string{"4"}
			So(Difference(slice1, slice2), ShouldResemble, result)
		})
	})

	Convey("Given the datasets", t, func() {
		data := [][]string{
			[]string{"id", "type", "value"},
			[]string{"111", "A", "2"},
			[]string{"111", "B", "1"},
			[]string{"111", "A", "1"},
			[]string{"222", "C", "2"},
			[]string{"222", "B", "1"},
			[]string{"222", "C", "3"},
		}

		x := []string{"id", "type"}
		y := "value"

		Convey("If I group by id and type, the value of the first array is [111, A, 3]", func() {
			result := Aggregate(x, y, data)
			So(result[1], ShouldResemble, []string{"111", "A", "3"})
		})
	})
}
