// main func using tempconv package
package main

import (
	"fmt"

	"github.com/unkcpz/gopl-exer/ch2/tempconv/tempconv"
)

func main() {
	boilingC := tempconv.BoilingC
	freezingC := tempconv.FreezingC
	fmt.Printf("Boiling Point is %v or %v\n", boilingC,
		tempconv.CToF(boilingC))
	fmt.Printf("Freezing Point is %v or %v\n", freezingC,
		tempconv.CToF(freezingC))
}
