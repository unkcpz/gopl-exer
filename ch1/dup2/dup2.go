// inputs from files give by cmd arguments, count the
// duplicate lines.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	counters := make(map[string]int)
	if len(files) == 0 {
		countLines(os.Stdin, counters)
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v, %v", f, err)
			}
			countLines(f, counters)
			f.Close()
		}
	}
	for line, n := range counters {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLines(f *os.File, counters map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counters[scanner.Text()]++
	}
}
