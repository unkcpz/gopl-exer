// Modify dup2 to print th enames of all files in which
// each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counters := make(map[string]int)
	foundIn := make(map[string][]string)
	args := os.Args[1:]
	if len(args) == 0 {
		countLines(os.Stdin, counters, foundIn)
	} else {
		for _, arg := range args {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex dup: %v, %v", f, err)
			}
			countLines(f, counters, foundIn)
		}
	}

	for line, n := range counters {
		if n > 1 {
			fmt.Printf("%v\t%d\t%s\n", foundIn[line], n, line)
		}
	}
}

func in(s string, sl []string) bool {
	for _, i := range sl {
		if s == i {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, fi map[string][]string) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		counts[line]++
		if !in(f.Name(), fi[line]) {
			fi[line] = append(fi[line], f.Name())
		}
	}
}
