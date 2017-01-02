package tools

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// mean function
func mean(x []float64) float64 {
	out := 0.00
	n := len(x)
	for i := 0; i < n; i++ {
		out = out + x[i]
	}
	out = out / float64(n)
	return out
}

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

		x := []string{"id", "day", "type"}
		y := "value"
		c := make(chan [][]string)

		Convey("If I group by id and type, the value of the first array is [111, A, 3]", func() {
			result := Aggregate(x, y, data)
			So(result[1], ShouldResemble, []string{"111", "A", "3"})
		})

		Convey("With Agreggate2, the value of the first array is [id type value]", func() {
			go Aggregate2(x, y, data, c)
			result := <-c
			So(result[0], ShouldResemble, []string{"id", "type", "value"})
		})

		Convey("With Pivot, the value of the first array is [0 1 5]", func() {
			go Aggregate2(x, y, data, c)
			result := <-c
			result2, result1 := Pivot("id", "type", result)
			fmt.Println(result2)
			So(result1[1], ShouldResemble, []float64{0, 1, 5})
		})

	})

	Convey("Given the following points", t, func() {
		x := [][]float64{
			[]float64{3, 2.5, 3.5, 1, 4.5},
			[]float64{1, 2, 0, 1, 1.6},
		}

		Convey("The mean of each raw of [[3 2.5 3.5 1 4.5] [1 2 0 1 1.6]] is [2.9 1.12]", func() {
			out := Apply(x, 1, mean)
			So(out, ShouldResemble, []float64{2.9, 1.1199999999999999})
		})

		Convey("The mean of each column of [[3 2.5 3.5 1 4.5] [1 2 0 1 1.6]] is [2 2.25 1.75 3.05]", func() {
			out := Apply(x, 2, mean)
			So(out, ShouldResemble, []float64{2, 2.25, 1.75, 1, 3.05})
		})

		order := Order(x[0], true)
		Convey("The decreasing order of [3 2.5 3.5 1 4.5] is [5 3 1 2 4]", func() {
			So(order, ShouldResemble, []int{5, 3, 1, 2, 4})
		})

		Convey("The decreasing sort of [3 2.5 3.5 1 4.5] is [4.5 3.5 3 2.5 1]", func() {
			y := [][]float64{
				[]float64{1, 1, 1},
				[]float64{2, 2, 2},
				[]float64{3, 3, 3},
				[]float64{4, 4, 4},
				[]float64{5, 5, 5},
			}
			sort := Sort(y, order)
			So(sort[0], ShouldResemble, []float64{5, 5, 5})
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
