// revarray reverses a 5 element array.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const l = 5

func main() {
	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints [l]int
		for i, s := range strings.Fields(input.Text()) {
			if i < len(ints) {
				x, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue outer
				}
				ints[i] = int(x)
			}
		}
		reverse(&ints)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

// reverse reverses an array of 5 ints in place.
func reverse(s *[l]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
