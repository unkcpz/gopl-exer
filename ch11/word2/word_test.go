package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
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
}
