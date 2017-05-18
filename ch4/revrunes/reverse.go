package reverse

import (
	"unicode/utf8"
)

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func ReverseRunes(s []byte) {
	reverse(s)
	i := 0
	for j, b := range s {
		if utf8.RuneStart(b) {
			reverse(s[i : j+1])
			i = j + 1
		}
	}
}
