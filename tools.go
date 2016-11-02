package tools

import (
	"fmt"
	"strconv"
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
		if encountered[x[v]] == true {
		} else {
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
	fmt.Println(out)
	return out
}

func Pivot(datasets [][]string) [][]string {

}
