// print a spinner when calculating fab goroutine
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	fmt.Printf("\rFibonacci(%d) = %d\n", 45, fab(45))
}

func spinner(delay time.Duration) {
	for {
		for _, c := range `\|/-` {
			fmt.Printf("\r%c", c)
			time.Sleep(delay)
		}
	}
}

func fab(x int) int {
	if x < 2 {
		return x
	}
	return fab(x-1) + fab(x-2)
}
