package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s string
	sep := ""

	start := time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("%s: %.8fs\n", s, time.Since(start).Seconds())

	start = time.Now()
	s = strings.Join(os.Args[1:], " ")
	fmt.Printf("%s: %.8fs\n", s, time.Since(start).Seconds())

}
