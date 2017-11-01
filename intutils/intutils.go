package intutils

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/entropyx/tools/strutils"
)

// GroupDigits groups each n digits of a number from right to left. Use sep as the seperator for each group.
func GroupDigits(number int, sep string, n int) string {
	str := strconv.Itoa(number)
	return strutils.GroupDigits(str, sep, n)
}

// RandInt returns a random integer between min and max
func RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

func Test() {

}
