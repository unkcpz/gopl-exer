// Rotate int slice in a single pass
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = rotate(s, 2)
	fmt.Println(s) // [3, 4, 5, 1, 2]
}

func rotate(s []int, n int) []int {
	st := make([]int, len(s))
	for i := 0; i < n; i++ {
		st[len(s)-n+i] = s[i]
	}
	for i := n; i < len(s); i++ {
		st[i-n] = s[i]
	}
	return st
}
