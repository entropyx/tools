package tools

import (
	"fmt"
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
			[]string{"id", "day", "type", "value"},
			[]string{"111", "M", "A", "2"},
			[]string{"111", "M", "B", "1"},
			[]string{"111", "M", "A", "1"},
			[]string{"222", "T", "C", "2"},
			[]string{"222", "W", "B", "1"},
			[]string{"222", "T", "C", "3"},
		}

		x := []string{"id", "day", "type"}
		y := "value"

		Convey("If I group by id and type, the value of the first array is [111, A, 3]", func() {
			result := Aggregate(x, y, data)
			So(result[1], ShouldResemble, []string{"111", "M", "A", "3"})
		})

		Convey("With Agreggate2, the value of the first array is [111, A, 3]", func() {
			result := Aggregate2(x, y, data)
			So(result[1], ShouldResemble, []string{"111", "M", "A", "3"})
		})

		Convey("If I pivot by id and day, the value of the first array is [111, M, 2, 1]", func() {
			data2 := Aggregate(x, y, data)
			result := Pivot([]string{"id", "day"}, "type", data2)
			So(result[1], ShouldResemble, []string{"111", "M", "3", "1", "0"})
		})

		Convey("With Pivot2, the value of the first array is [0 1 5]", func() {
			data2 := Aggregate2(x, y, data)
			result2, result1 := Pivot2("id", "type", data2)
			fmt.Println(result2)
			So(result1[1], ShouldResemble, []float64{0, 1, 5})
		})

	})

	Convey("Given two points", t, func() {

		x := []float64{1, 0, 0, 3, 4, 42.1, 3}
		y := []float64{1, 2, 3.4, 0, 0, 8.2, 0}

		Convey("The distance between x and y is ", func() {
			distance := Dist(x, y)
			So(distance, ShouldResemble, 34.62325807892724)
		})

	})
}
