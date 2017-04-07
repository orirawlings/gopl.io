// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//	$ comma 1 12 123 -1234.6509 -12345 +123456 1234567890
// 	1
// 	12
// 	123
// 	-1,234.6509
//	-12,345
//	+123,456
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a decimal string.
func comma(s string) string {
	var b bytes.Buffer
	// Optional sign
	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		s = s[1:]
	}
	// Optional fractional portion
	var frac string
	if i := strings.IndexByte(s, '.'); i > -1 {
		frac = s[i:]
		s = s[:i]
	}
	p := len(s) % 3
	if p == 0 && len(s) > 0 {
		p = 3
	}
	b.WriteString(s[:p])
	for i := p + 3; i <= len(s); i += 3 {
		b.WriteString(",")
		b.WriteString(s[i-3 : i])
	}
	b.WriteString(frac)
	return b.String()
}
