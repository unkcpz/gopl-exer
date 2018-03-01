// Reverse use an array pointer instead of a slice
package main

import "fmt"

func main() {
	s := [5]int{1, 2, 3, 4, 5}
	reverse(&s)
	fmt.Println(s)
}

func reverse(ps *[5]int) {
	for i, j := 0, len(ps)-1; i < j; i, j = i+1, j-1 {
		// (*ps)[i], ps[j] = ps[j], ps[i] // ok!
		ps[i], ps[j] = ps[j], ps[i]
	}
}
