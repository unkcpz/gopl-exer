package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("random") {
		t.Error(`IsPalindrome("random") = true`)
	}
}

func TestChinesePalindrome(t *testing.T) {
	if !IsPalindrome("上海自来水来自海上") {
		t.Error(`IsPalindrome("上海自来水来自海上") = false`)
	}
}

func TestCanalPalindrom(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrom(%q) = false`, input)
	}
}
