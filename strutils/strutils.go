package strutils

// GroupDigits groups each n digits of a number from right to left. Use sep as the seperator for each group.
func GroupDigits(str, sep string, n int) string {
	var groupedString []byte
	sepb := []byte(sep)
	counter := 0
	offset := len(str) % n
	for i := 0; i < len(str); i++ {
		groupedString = append(groupedString, str[i])
		if i+1 < offset {
			continue
		}

		if offset == 0 && counter == 0 {
			counter++
		}

		if (i+1 == offset || (counter > 0 && counter%n == 0)) && i+1 < len(str) {
			groupedString = append(groupedString, sepb...)
		}
		counter++
	}
	return string(groupedString)
}
