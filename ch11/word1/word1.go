// word provides utilities for word game
package word

// IsPalindrome check word is palindrome
func IsPalindrome(s string) bool {
	for i, _ := range s {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
