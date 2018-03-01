// in-place technology to eliminate adjacent duplicates in a []sting slice
package main

import "fmt"

func main() {
	s := []string{"w", "z", "z", "y", "y", "y", "qw", "po", "qw", "w"}
	s = rad(s)
	fmt.Printf("%q\n", s) // ["w", "z", "y", "qw", "po", "qw", "w"]
}

func rad(ss []string) []string {
	idx := 0
	for i := 1; i < len(ss); i++ {
		if ss[i] != ss[i-1] {
			ss[idx] = ss[i]
			idx++
		}
	}
	return ss[:idx]
}
