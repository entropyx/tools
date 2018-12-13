package strutils

import (
	"strconv"
	"strings"
)

const (
	NormalizeEurope = iota
	NormalizeAmericanBritain
)

var euList = []string{
	"EUR", "GBP",
}

func normalizeEurope(old string) string {
	count := strings.Count(old, ".")
	s := strings.Replace(old, ",", ".", -1)
	return strings.Replace(s, ".", "", count)

}
func normalizeAmerican(old string) string {
	return strings.Replace(old, ",", "", -1)
}

// ParseCurrency determines locale and converts it to float64
func ParseCurrency(code string, fs string) (float64, error) {
	for _, c := range euList {
		if code == c {
			return strconv.ParseFloat(normalizeEurope(fs), 64)
		}
	}
	return strconv.ParseFloat(normalizeAmerican(fs), 64)
}
