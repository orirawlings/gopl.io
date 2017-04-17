// Eliminate adjacent duplicate strings in string slice
package main

import "fmt"

func dedup(strings []string) []string {
	var i int
	var last string
	for _, s := range strings {
		if i == 0 || s != last {
			strings[i] = s
			i++
		}
		last = s
	}
	return strings[:i]
}

func main() {
	fmt.Printf("%q\n", dedup([]string{}))
	fmt.Printf("%q\n", dedup([]string{""}))
	fmt.Printf("%q\n", dedup([]string{"", ""}))
	fmt.Printf("%q\n", dedup([]string{"a", "b"}))
	fmt.Printf("%q\n", dedup([]string{"a", "a"}))
	fmt.Printf("%q\n", dedup([]string{"a", "a", "b"}))
	fmt.Printf("%q\n", dedup([]string{"a", "b", "b"}))
	fmt.Printf("%q\n", dedup([]string{"a", "b", "c", "b"}))
	fmt.Printf("%q\n", dedup([]string{"a", "b", "b", "b", "c"}))
	fmt.Printf("%q\n", dedup([]string{"car", "plane", "plane", "boat"}))
}
