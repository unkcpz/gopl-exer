// Using io/ioutil ReadFile read file contents
// Then count the duplicate lines
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counters := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v, %v", filename, err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counters[line]++
		}
	}
	for line, n := range counters {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
