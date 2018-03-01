// nonempty returns a slice holding only the non-empty strings
package main

import "fmt"

func main() {
	fmt.Printf("%q\n", nonempty([]string{"y", "", "notempty"})) // ["y", "notempry"]
}

func nonempty(ss []string) []string {
	i := 0
	for _, s := range ss {
		if s != "" {
			ss[i] = s
			i++
		}
	}
	return ss[:i]
}
