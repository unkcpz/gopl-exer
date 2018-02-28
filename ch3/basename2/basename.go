// Getting basename using string.LastIndex()
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}

func basename(s string) string {
	slashIdx := strings.LastIndex(s, "/")
	s = s[slashIdx+1:]
	dotIdx := strings.LastIndex(s, ".")
	if dotIdx := strings.LastIndex(s, "."); dotIdx >= 0 {
		s = s[:dotIdx]
	}
	return s
}
