// Ftoc prints two Fahrenheit-to-Celsius conversions
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("freezing: %gF = %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("boiling: %gF = %gC\n", boilingF, fToC(boilingF))
}

func fToC(ft float64) float64 {
	ct := (ft - 32) * 5 / 9
	return ct
}
