package strutils

import (
	"math/rand"
	"time"

	"github.com/entropyx/tools/intutils"
)

//Rand generates a random string with l length
func Rand(l int) string {
	var s string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l; i++ {
		n := intutils.RandUint(33, 126)
		s += string(n)
	}
	return s
}
