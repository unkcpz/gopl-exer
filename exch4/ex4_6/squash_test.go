package squash

import "testing"

func TestSquashSpace(t *testing.T) {
	var tests = []struct {
		input  []byte
		wanted []byte
	}{
		{[]byte("abc\t\n \n  ab"), []byte("abc ab")},
		{[]byte("abc\t\n \n  ab中文"), []byte("abc ab中文")},
	}

	for _, test := range tests {
		if got := string(SquashSpace(test.input)); got != string(test.wanted) {
			t.Errorf("SquashSpace(%q) = %q, want %q", test.input, got, test.wanted)
		}
	}
}
