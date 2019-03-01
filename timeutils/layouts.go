package timeutils

import (
	"strings"
	"time"
)

const (
	LayoutYYYY = "2006"
	LayoutYY   = "06"

	LayoutMMMM = "January"
	LayoutMMM  = "Jan"
	LayoutMM   = "01"

	LayoutEEE  = "Mon"
	LayoutEEEE = "Monday"

	LayoutD = "02"
)

var bigEndian = []string{
	LayoutYYYY, LayoutMM, LayoutD,
}

var littleEndian = []string{
	LayoutD, LayoutMM, LayoutYYYY,
}

var middleEndian = []string{
	LayoutMM, LayoutD, LayoutYYYY,
}

//BigEndian codes time in the format "year, month, day". Each component
//is separated by sep.
func BigEndian(t time.Time, sep string) string {
	layout := strings.Join(bigEndian, sep)
	return t.Format(layout)
}

//LittleEndian codes time in the format "day, month, year". Each
//component is separated by sep.
func LittleEndian(t time.Time, sep string) string {
	layout := strings.Join(littleEndian, sep)
	return t.Format(layout)
}

//MiddleEndian codes time in the format "day, month, year". Each
//component is separated by sep.
func MiddleEndian(t time.Time, sep string) string {
	layout := strings.Join(middleEndian, sep)
	return t.Format(layout)
}
