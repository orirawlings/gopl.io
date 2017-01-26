// Dup4 prints the count and text of lines that appear more than once
// in the input. And a list of all the files the line appears in. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fnames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts, fnames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, fnames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%q\t%v\n", n, line, fnames[line])
		}
	}
}

func contains(ss []string, n string) bool {
	for _, s := range ss {
		if s == n {
			return true
		}
	}
	return false
}

func countLines(f *os.File, fname string, counts map[string]int, fnames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !contains(fnames[line], fname) {
			fnames[line] = append(fnames[line], fname)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
