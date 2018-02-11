// Print the boiling temperature of water
package main

import "fmt"

const boilingF float64 = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("The boiling point of water is %.2fF or %.2fC\n", f, c)
}
