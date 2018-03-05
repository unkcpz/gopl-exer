// wordfreq counts frequency of each word in an input text file
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var freq = make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		freq[input.Text()]++
	}
	for words, fs := range freq {
		fmt.Printf("%d\t%s\n", fs, words)
	}
}
