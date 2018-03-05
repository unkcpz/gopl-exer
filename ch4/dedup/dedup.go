// Read stdin and print non-duplicated lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if !lines[text] {
			lines[text] = true
			fmt.Println(text)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
