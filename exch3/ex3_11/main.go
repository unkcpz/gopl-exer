// A non-recursive version of comma, using byte.Buffer
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("1234567890"))
}

func comma(s string) string {
	var buf bytes.Buffer
	pre := len(s) % 3
	buf.WriteString(s[:pre])
	for i := pre; i < len(s); i += 3 {
		if i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

func commaEnhance(s string) string {
	var buf bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	dot := strings.Index(s, ".")
	if dot > 0 {
		buf.WriteString(comma(s[:dot]) + s[dot:])
	} else {
		buf.WriteString(comma(s))
	}
	return buf.String()
}
