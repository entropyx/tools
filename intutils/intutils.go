package intutils

import (
	"strconv"

	"github.com/entropyx/tools/strutils"
)

// GroupDigits groups each n digits of a number from right to left. Use sep as the seperator for each group.
func GroupDigits(number int, sep string, n int) string {
	str := strconv.Itoa(number)
	return strutils.GroupDigits(str, sep, n)
}

func Test() {

}
