// sum function use variadic parameter
package main

import "fmt"

func main() {
	fmt.Println(sum())           //0
	fmt.Println(sum(1))          //1
	fmt.Println(sum(1, 2, 3, 4)) //10

	var s = []int{1, 2, 3}
	fmt.Println(sum(s...)) //6
}

func sum(xs ...int) int {
	s := 0
	for _, x := range xs {
		s += x
	}
	return s
}
