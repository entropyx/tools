package tools

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

func Aggregate(colgroupnames []string, colvaluename string, datasets [][]string) []int {
	var index []int
	//l1 := len(colgroupnames)
	l2 := len(datasets[0])
	//l3 := len(datasets)
	for _, col := range colgroupnames {
		for i := 0; i < l2; i++ {
			if col == datasets[0][i] {
				index = append(index, i)
			}
		}
	}
	return index
}
