// rotate rotates a slice left by a given number of positions
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Interactive test of rotate.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		rotate(ints[0], ints[1:])
		fmt.Printf("%v\n", ints[1:])
	}
	// NOTE: ignoring potential errors from input.Err()
}

// rotate rotates slice, s, left by, pos, positions
func rotate(pos int, s []int) {
	pos = pos % len(s)
	b := make([]int, pos)
	copy(b, s[:pos])
	copy(s, s[pos:])
	copy(s[len(s)-pos:], b)
}
