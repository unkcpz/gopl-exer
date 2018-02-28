package main

import "testing"

func TestAnagram(t *testing.T) {
	var tests = []struct {
		a, b   string
		wanted bool
	}{
		{"python", "typhon", true},
		{"banana", "abnana", true},
		{"Madam Curie", "Radium came", true},
		{"William Shakespeare", "I am a weakish speller", true},
		{"radom", "not same", false},
		{"falsetest", "falsetesthhh", false}, // coverage
		{"天气真好哈哈哈", "哈哈哈真气天好", true},
	}
	for _, test := range tests {
		if got := anagram(test.a, test.b); got != test.wanted {
			t.Errorf("anagram(%q, %q) == %t", test.a, test.b, got)
		}
	}
}
