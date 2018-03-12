// variadic version of strings.Join
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Join(" ", "hello", "world", ".")) // hello world.
	fmt.Println(Join(""))                         // ""
}

func Join(sep string, s ...string) string {
	return strings.Join(s, sep)
}
