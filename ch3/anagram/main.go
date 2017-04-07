// Compute whether two UTF-8 strings are anagrams. Program exits with 0 if anagrams, 1 otherwise.
//
// Usage:
// 	$ anagram <word1> <word2>
//
// Example:
// 	$ anagram silent listen
//
package main

import (
	"os"
)

func main() {
	if anagram(os.Args[1], os.Args[2]) {
		return
	}
	os.Exit(1)
}

func runefreq(s string) map[rune]int {
	r := make(map[rune]int)
	for _, c := range s {
		r[c] += 1
	}
	return r
}

func anagram(a, b string) bool {
	if a == b || len(a) != len(b) {
		return false
	}
	arunes, brunes := runefreq(a), runefreq(b)
	if len(arunes) != len(brunes) {
		return false
	}
	for r, af := range arunes {
		if bf := brunes[r]; bf != af {
			return false
		}
	}
	return true
}
