// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	s, sep := "", ""
	start := time.Now()
	for index, arg := range os.Args[1:] {
		s += sep + "(" + fmt.Sprint(index) + "," + arg + ")"
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%2fs elapsed", time.Since(start).Seconds())
}

//!-
