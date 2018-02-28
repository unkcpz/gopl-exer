package main

import "testing"

func TestComma(t *testing.T) {
	var tests = []struct {
		input  string
		wanted string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
		{"12345678", "12,345,678"},
	}

	for _, test := range tests {
		if got := comma(test.input); got != test.wanted {
			t.Errorf("comma(%s) = %s, want %s", test.input, got, test.wanted)
		}
	}
}

func TestCommaEnhance(t *testing.T) {
	var tests = []struct {
		input  string
		wanted string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
		{"12345678", "12,345,678"},
		{"-123456", "-123,456"},
		{"+1234567.542", "+1,234,567.542"},
		{"12345678.123", "12,345,678.123"},
		{"12345678.", "12,345,678."},
	}

	for _, test := range tests {
		if got := commaEnhance(test.input); got != test.wanted {
			t.Errorf("commaEnhance(%s) = %s, want %s", test.input, got, test.wanted)
		}
	}
}
