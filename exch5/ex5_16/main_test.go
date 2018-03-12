package main

import "testing"

func TestJoin(t *testing.T) {
	var tests = []struct {
		input  []string
		wanted string
	}{
		{[]string{" ", "oh", "my", "haha"}, "oh my haha"},
	}
	for _, test := range tests {
		if got := Join(test.input[0], test.input[1:]...); got != test.wanted {
			t.Errorf("Join(%q) = %q want %q", test.input, got, test.wanted)
		}
	}
}
