package popcount

import "testing"

func TestPopCount(t *testing.T) {
	var tests = []struct {
		input  uint64
		wanted int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{4, 1},
		{1024, 1},
		{0xfffff, 20},
		{255, 8},
	}

	for _, test := range tests {
		if got := PopCount(test.input); got != test.wanted {
			t.Errorf("PopCount(%v) = %d, want %d", test.input, got, test.wanted)
		}
	}

	for _, test := range tests {
		if got := PopCountLoop(test.input); got != test.wanted {
			t.Errorf("PopCountLoop(%v) = %d, want %d", test.input, got, test.wanted)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0xffffffff)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(0xffffffff)
	}
}
