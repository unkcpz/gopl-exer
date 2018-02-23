// Use popcount package
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/unkcpz/gopl-exer/ch2/popcount/popcount"
)

func main() {
	i, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal("input wrong")
	}
	fmt.Println(popcount.PopCount(uint64(i)))
}
