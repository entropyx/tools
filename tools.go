package tools

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type aux struct {
	ids    []string
	values []float64
}

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
		if !encountered[x[v]] {
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

func Compact(x [][]string, index1 []int, index2 int) map[string]map[string]float64 {
	encountered := make(map[string]map[string]bool)
	result := make(map[string]map[string]float64)
	for v := range x {
		id1 := x[v][index1[0]]
		id2 := x[v][index1[1]]
		value := x[v][index2]
		if result[id1] == nil {
			result[id1] = map[string]float64{}
		}

		if encountered[id1] == nil {
			encountered[id1] = map[string]bool{}
		}

		if encountered[id1][id2] == false {
			encountered[id1][id2] = true
			a, _ := strconv.ParseFloat(value, 64)
			result[id1][id2] = a
		} else {
			a, _ := strconv.ParseFloat(value, 64)
			result[id1][id2] = result[id1][id2] + a
		}
	}
	return result
}

func Aggregate2(colgroupnames []string, colvaluename string, datasets [][]string, c chan [][]string) {
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
	out = append(out, datasets[0])
	y := Compact(datasets[1:], index1, index2)

	for key1, map2 := range y {
		for key2, value := range map2 {
			out = append(out, append([]string{key1, key2, fmt.Sprintf("%v", value)}))
		}
	}
	c <- out
}

func Pivot(colpivotname string, colvaluename string, datasets [][]string) (u2 []string, pivot [][]float64) {
	var index1, index2, index3 int
	var aux1, aux2 []string

	for i := 0; i < len(datasets[0]); i++ {
		if colvaluename == datasets[0][i] {
			index1 = i
		} else if "value" == datasets[0][i] {
			index2 = i
		} else if colpivotname == datasets[0][i] {
			index3 = i
		}
	}

	for i := 1; i < len(datasets); i++ {
		aux1 = append(aux1, datasets[i][index3])
		aux2 = append(aux2, datasets[i][index1])
	}
	u1 := Unique(aux1)
	u2 = Unique(aux2)
	result := make(map[string]map[string]float64)

	for i := 1; i < len(datasets); i++ {
		id1 := datasets[i][index3]
		id2 := datasets[i][index1]
		value, _ := strconv.ParseFloat(datasets[i][index2], 64)
		if result[id1] == nil {
			result[id1] = map[string]float64{}
		}
		result[id1][id2] = value
	}

	for i := range u1 {
		row := []float64{}
		for j := range u2 {
			row = append(row, []float64{result[u1[i]][u2[j]]}...)
		}
		pivot = append(pivot, row)
	}
	return
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

func TimeMinusDays(days int) string {
	layout := "2006-01-02T15:04:05.000Z"
	t := time.Now().AddDate(0, 0, days)
	fmt.Println(t)
	mongoTime := t.Format(layout)
	return mongoTime
}

func NanoSecondsToMongoTime(ns int64) string {
	layout := "2006-01-02T15:04:05.000Z"
	fmt.Println(ns)
	t := fmt.Sprintf("%d", ns)
	mongoString, _ := time.Parse(layout, t)
	return mongoString.String()
}
