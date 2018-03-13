// pipeline use chan as pipeline to get square of naturals
package main

import "fmt"

func main() {
	natruals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			natruals <- x
		}
	}()

	// Squares
	go func() {
		for {
			x := <-natruals
			squares <- x * x
		}
	}()
	// Printer
	for {
		fmt.Println(<-squares)
	}
}
