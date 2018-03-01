// Counts the number of bits that are different in two SHA256 hashes
package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func main() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	// fmt.Println(popCount(10))
	c1 := sha256.Sum256([]byte("x")) // c1 -> [32]uint8
	c2 := sha256.Sum256([]byte("X"))

	sum := 0
	for i, _ := range c1 {
		sum += popCount(c1[i] &^ c2[i])
	}
	fmt.Printf("c1:\t%08b\nc2:\t%08b\n%d\n", c1, c2, sum)
}

func popCount(x uint8) int {
	return int(pc[x])
}
