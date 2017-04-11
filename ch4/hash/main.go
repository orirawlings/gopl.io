// Computes the SHA256 (default), SHA384, or SHA512 hash of stdin.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	a := flag.String("a", "sha256", "The hash algorithm to use, options include (sha256, sha384, sha512)")
	flag.Parse()
	h := sha256.New()
	switch *a {
	case "sha384":
		h = sha512.New384()
	case "sha512":
		h = sha512.New()
	}
	if _, err := io.Copy(h, os.Stdin); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", h.Sum(nil))
}
