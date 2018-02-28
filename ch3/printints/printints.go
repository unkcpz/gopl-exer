// printints like fmt.Sprint(value) but add comma
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", intsToString([]int{1, 2, 3})) // [1, 2, 3]
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteRune('左')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
