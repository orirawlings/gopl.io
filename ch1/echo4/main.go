// Echo3 prints its name and command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	fmt.Println(strings.Join(os.Args, " "))
}

//!-
