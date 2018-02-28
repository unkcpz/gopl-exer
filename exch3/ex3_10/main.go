// A non-recursive version of comma, using byte.Buffer
package main

import (
	"bytes"
	"fmt"
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
