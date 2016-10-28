package tools

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

func unique(x []string) []string {
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

func removeElement(element string, array []string) []string {
	var out []string
	l1 := len(array)
	for i := 0; i < l1; i++ {
		if array[i] != element {
			out = append(out, array[i])
		}
	}
	return out
}
