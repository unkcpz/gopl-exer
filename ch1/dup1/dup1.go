// Read inputs from stdin, write and count lines appear more than one time
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// read from os.Stdin
	counters := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		counters[scanner.Text()]++
	}

	// write the results
	for line, n := range counters {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
