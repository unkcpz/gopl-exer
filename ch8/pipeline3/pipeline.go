// pipeline use chan to print squares of naturals
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printfun(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printfun(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
