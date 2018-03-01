package reverse

import (
	"testing"
)

func TestReverseUTF8(t *testing.T) {
	input := []byte("中文字")
	wanted := []byte("字文中")
	if got := ReverseUTF8(input); string(got) != string(wanted) {
		t.Errorf("RevereseUTF8(%q) = %q, want %q", input, got, wanted)
	}
}
