package tools

import (
  "time"
  "fmt"
	"math"
	"strconv"
	"strings"
)

// Test if given one element is subset of an array.
func Subset(first, second []string) bool {
	set := make(map[string]int)
	for _, value := range second {
		set[value]++
	}
	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

// Remove duplicate elements of the given array.
func Unique(x []string) []string {
	encountered := map[string]bool{}
	result := []string{}
	for v := range x {
		if encountered[x[v]] != true {
			encountered[x[v]] = true
			result = append(result, x[v])
		}
	}
	return result
}

// Remove array elements given element that I want remove.
func RemoveElement(element string, array []string) []string {
	var out []string
	l1 := len(array)
	for i := 0; i < l1; i++ {
		if array[i] != element {
			out = append(out, array[i])
		}
	}
	return out
}

// Remove the i'ths index of array of array of arrays.
func RemoveArray(index int, arrays [][]string) [][]string {
	arrays = append(arrays[:index], arrays[index+1:]...)
	return arrays
}

// Difference returns difference between two slices of strings
func Difference(slice1 []string, slice2 []string) []string {
	var difference []string

	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				difference = append(difference, s1)
			}
		}
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return difference
}

func Aggregate(colgroupnames []string, colvaluename string, datasets [][]string) [][]string {
	var index1 []int
	var index2 int
	var ss bool
	var out [][]string
	l1 := len(datasets[0])
	for _, col := range colgroupnames {
		for i := 0; i < l1; i++ {
			if col == datasets[0][i] {
				index1 = append(index1, i)
			}
		}
	}
	out = append(out, datasets[0])
	for i := 0; i < l1; i++ {
		if colvaluename == datasets[0][i] {
			index2 = i
			break
		}
	}
	i := 1
L:
	for i < len(datasets) {
		var row []string
		for _, j1 := range index1 {
			row = append(row, datasets[i][j1])
		}
		for j := 0; j < len(out); j++ {
			ss = Subset(row, out[j])
			if ss {
				a, err1 := strconv.ParseFloat(datasets[i][index2], 64)
				b, err2 := strconv.ParseFloat(out[j][index2], 64)
				if err1 != nil {
					fmt.Println(err1)
				} else if err2 != nil {
					fmt.Println(err2)
				}
				c := a + b
				out[j][index2] = fmt.Sprintf("%v", c)
				i = i + 1
				goto L
			}
		}
		out = append(out, datasets[i])
		i = i + 1
	}
	return out
}

type aux struct {
	ids    []string
	values []float64
}

func Aggregate2(colgroupnames []string, colvaluename string, datasets [][]string) [][]string {
	var out [][]string
	var index1 []int
	var index2 int

	l1 := len(datasets[0])
	for _, col := range colgroupnames {
		for i := 0; i < l1; i++ {
			if col == datasets[0][i] {
				index1 = append(index1, i)
			}
		}
	}

	for i := 0; i < l1; i++ {
		if colvaluename == datasets[0][i] {
			index2 = i
			break
		}
	}

	agg := new(aux)

	for i := 1; i < len(datasets); i++ {
		var row []string
		for _, j1 := range index1 {
			row = append(row, datasets[i][j1])
		}
		agg.ids = append(agg.ids, strings.Join(row, ","))
		a, _ := strconv.ParseFloat(datasets[i][index2], 64)
		agg.values = append(agg.values, a)
	}

	out = append(out, datasets[0])
	uniqueids := Unique(agg.ids)

	for _, i := range uniqueids {
		total := 0.00
		for j := 0; j < len(agg.ids); j++ {
			if i == agg.ids[j] {
				total = total + agg.values[j]
			}
		}
		out = append(out, append(strings.Split(i, ","), fmt.Sprintf("%v", total)))
	}

	return out
}

func Pivot(colpivotnames []string, colvaluename string, datasets [][]string) [][]string {
	var pivot [][]string
	var index1 int
	var index2 int
	var index3 []int
	var aux []string
	var s bool
	var ss bool
	var acum []int
	pivot = append(pivot, datasets[0][:len(colpivotnames)])
	for i := 0; i < len(datasets[0]); i++ {
		if colvaluename == datasets[0][i] {
			index1 = i
		} else if "value" == datasets[0][i] {
			index2 = i
		}
	}
	for _, col := range colpivotnames {
		for i := 0; i < len(colpivotnames); i++ {
			if col == datasets[0][i] {
				index3 = append(index3, i)
			}
		}
	}
	for i := 1; i < len(datasets); i++ {
		aux = append(aux, datasets[i][index1])
	}
	u := Unique(aux)
	pivot[0] = append(pivot[0], u...)
	i2 := 0
	for i := 1; i < len(datasets); i++ {
		s = false
		for _, r := range acum {
			if i == r {
				s = true
				break
			}
		}
		if !s {
			i2 = i2 + 1
			var row []string
			for _, j1 := range index3 {
				row = append(row, datasets[i][j1])
			}
			pivot = append(pivot, row)
			for _, j3 := range u {
				for k := 1; k < len(datasets); k++ {
					ss = Subset(append(row, j3), datasets[k])
					if ss {
						pivot[i2] = append(pivot[i2], datasets[k][index2])
						acum = append(acum, k)
						break
					}
				}
				if ss != true {
					pivot[i2] = append(pivot[i2], "0")
				}
			}
		}
	}
	return pivot
}

// Function Dist compute distance between two points.
func Dist(point1, point2 []float64) float64 {
	var dist float64
	if len(point1) != len(point2) {
		fmt.Println("Vectors of different length!")
	} else {
		a := 0.00
		for i := 0; i < len(point1); i++ {
			a = a + math.Pow(point1[i]-point2[i], 2)
		}
		dist = math.Sqrt(a)
	}
	return dist
}


func TimeMinusDays(days int) (string) {
  layout := "2006-01-02T15:04:05.000Z"
  t := time.Now().AddDate(0, 0, days)
  fmt.Println(t)
  mongoTime := t.Format(layout)
  return mongoTime
}

func NanoSecondsToMongoTime(ns int64) (string) {
  layout := "2006-01-02T15:04:05.000Z"
  fmt.Println(ns)
  t := fmt.Sprintf("%d", ns)
  mongoString, _ := time.Parse(layout, t)
  return mongoString.String()
}
