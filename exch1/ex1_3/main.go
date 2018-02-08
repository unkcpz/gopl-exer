// The three echo function to be benchmarked
package main

import (
	"strings"
)

func main() {
	// args := os.Args[:]
	// echo1(args)
	// echo2(args)
	// echo3(args)
}

func echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	// fmt.Println(s)
}

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

func echo3(args []string) {
	strings.Join(args[1:], " ")
}
