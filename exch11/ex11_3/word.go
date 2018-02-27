// word package contains utilities for word game
package word

import "unicode"

func IsPalindrome(s string) bool {
	var runes []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			runes = append(runes, unicode.ToLower(r))
		}
	}

	for i, _ := range runes {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}
