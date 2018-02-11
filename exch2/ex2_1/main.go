// Using package tempconv
package main

import (
	"fmt"

	"github.com/unkcpz/gopl-exer/exch2/ex2_1/tempconv"
)

func main() {
	fmt.Printf("Absolute zero temperature is: %v, %v, %v\n",
		tempconv.KToC(tempconv.AbsoluteZeroK),
		tempconv.KToF(tempconv.AbsoluteZeroK),
		tempconv.AbsoluteZeroK)
}
