// Reverse the []byte slice represents a URF-8 string
// in-place
package reverse

import "unicode/utf8"

// algo: reverse each rune by bytes
// End with reverse the whole slice
func ReverseUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i : i+size])
		i += size
	}
	rev(b)
	return b
}

func rev(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
