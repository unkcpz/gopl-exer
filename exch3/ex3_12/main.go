// Reports whether two strings are anagrams of each other
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(anagram("ba", "ab")) // true
}

func anagram(a, b string) bool {
	aRuneCount := make(map[rune]int)
	bRuneCount := make(map[rune]int)
	for _, c := range a {
		if unicode.IsLetter(c) {
			aRuneCount[unicode.ToLower(c)]++
		}
	}

	for _, c := range b {
		if unicode.IsLetter(c) {
			bRuneCount[unicode.ToLower(c)]++
		}
	}

	for k, v := range aRuneCount {
		if bRuneCount[k] != v {
			return false
		}
	}

	for k, v := range bRuneCount {
		if aRuneCount[k] != v {
			return false
		}
	}

	return true
}
