// Scan standard input for words until EOF and then print frequency of each word
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freqs := make(map[string]int)
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		freqs[s.Text()]++
	}
	fmt.Println("word\tfreq")
	for word, freq := range freqs {
		fmt.Printf("%q\t%d\n", word, freq)
	}
	if err := s.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %+v", err)
	}
}
