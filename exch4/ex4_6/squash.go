// An in-place function  squashes adjacent Unicode spaces into a
// single ASCII space
package squash

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func SquashSpace(s []byte) []byte {
	idx := 0
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRune(s[i:])
		rNext, _ := utf8.DecodeRune(s[i+size:]) // out of range?? No but take care
		fmt.Println(len(s), s[i+size:])
		if unicode.IsSpace(r) && unicode.IsSpace(rNext) {
			i += size
			continue
		}
		utf8.EncodeRune(s[idx:], r)
		i += size
		idx += size
	}
	return s[:idx]
}
