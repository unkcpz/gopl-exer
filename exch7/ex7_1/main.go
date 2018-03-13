// Implement counters for words and for lines.
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	return int(*c), scanner.Err()
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	for scanner.Scan() {
		*c++
	}
	return int(*c), scanner.Err()
}

func main() {
	var wc WordCounter
	fmt.Fprintf(&wc, "format hello")
	fmt.Println(wc) // 2

	var lc LineCounter
	fmt.Fprintf(&lc, "first Line\nSecond Line\nThird Line\n\n")
	fmt.Println(lc) // 4
}
