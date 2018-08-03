// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var s string
	sep := " "
	start := time.Now()

	s += "Command:" + sep + os.Args[0] + "\n"
	s += "Arguments:"

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)
	fmt.Printf("%2fs elapsed", time.Since(start).Seconds())
}

//!-
