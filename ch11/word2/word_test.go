package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{" ", true},
		{"上海自来水来自海上", true},
		{"detartrated", true},
		{"kayak", true},
		{"random", false},
		{"été", true},
		{"Été le bar arabe l'été", true},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf(`IsPalindrome(%q) = %t`, test.input, got)
		}
	}

	for _, test := range tests {
		if got := IsPalindromeFast(test.input); got != test.want {
			t.Errorf(`IsPalindrome(%q) = %t`, test.input, got)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("Été le bar arabe l'été")
	}
}

func BenchmarkIsPalindromeFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindromeFast("Été le bar arabe l'été")
	}
}
