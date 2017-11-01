package strutils

import (
	"unicode"
	"unicode/utf8"
)

// Copy makes a copy of a string
func Copy(str string) string {
	b := str
	b2 := make([]byte, len(b))
	copy(b2, b)
	return string(b2)
}

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

type buffer struct {
	r         []byte
	runeBytes [utf8.UTFMax]byte
}

func (b *buffer) write(r rune) {
	if r < utf8.RuneSelf {
		b.r = append(b.r, byte(r))
		return
	}
	n := utf8.EncodeRune(b.runeBytes[0:], r)
	b.r = append(b.r, b.runeBytes[0:n]...)
}

func (b *buffer) indent() {
	if len(b.r) > 0 {
		b.r = append(b.r, '_')
	}
}

func Underscore(s string) string {
	b := buffer{
		r: make([]byte, 0, len(s)),
	}
	var m rune
	var w bool
	for _, ch := range s {
		if unicode.IsUpper(ch) {
			if m != 0 {
				if !w {
					b.indent()
					w = true
				}
				b.write(m)
			}
			m = unicode.ToLower(ch)
		} else {
			if m != 0 {
				b.indent()
				b.write(m)
				m = 0
				w = false
			}
			b.write(ch)
		}
	}
	if m != 0 {
		if !w {
			b.indent()
		}
		b.write(m)
	}
	return string(b.r)
}
