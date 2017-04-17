// Squash multiple adjacent spaces into single space within UTF-8 byte slice
package main

import (
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func trimspace(text []byte) []byte {
	var j int
	for i := 0; i < len(text); {
		r, size := utf8.DecodeRune(text[i:])
		if unicode.IsSpace(r) && i > 0 {
			if last, lsize := utf8.DecodeLastRune(text[:j]); unicode.IsSpace(last) {
				j -= lsize
				j += utf8.EncodeRune(text[j:], ' ')
			} else {
				j += utf8.EncodeRune(text[j:], r)
			}
		} else {
			j += utf8.EncodeRune(text[j:], r)
		}
		i += size
	}
	return text[:j]
}

func main() {
	for _, s := range os.Args[1:] {
		fmt.Printf("%s\n", trimspace([]byte(s)))
	}
}
