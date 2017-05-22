// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	control = iota
	digit
	graphic
	letter
	lower
	mark
	number
	print
	punct
	space
	symbol
	title
	upper
	invalid
)

var (
	categories = [...]string{
		"control",
		"digit",
		"graphic",
		"letter",
		"lower",
		"mark",
		"number",
		"print",
		"punct",
		"space",
		"symbol",
		"title",
		"upper",
		"invalid",
	}
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	var catecounts [invalid + 1]int // counts of Unicode characters in each category

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			catecounts[invalid]++
			continue
		}
		if unicode.IsControl(r) {
			catecounts[control]++
		}
		if unicode.IsDigit(r) {
			catecounts[digit]++
		}
		if unicode.IsGraphic(r) {
			catecounts[graphic]++
		}
		if unicode.IsLetter(r) {
			catecounts[letter]++
		}
		if unicode.IsLower(r) {
			catecounts[lower]++
		}
		if unicode.IsMark(r) {
			catecounts[mark]++
		}
		if unicode.IsNumber(r) {
			catecounts[number]++
		}
		if unicode.IsPrint(r) {
			catecounts[print]++
		}
		if unicode.IsPunct(r) {
			catecounts[punct]++
		}
		if unicode.IsSpace(r) {
			catecounts[space]++
		}
		if unicode.IsSymbol(r) {
			catecounts[symbol]++
		}
		if unicode.IsTitle(r) {
			catecounts[title]++
		}
		if unicode.IsUpper(r) {
			catecounts[upper]++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\ncategory\tcount\n")
	for i, count := range catecounts {
		if count > 0 {
			fmt.Printf("%s\t%d\n", categories[i], count)
		}
	}
}
