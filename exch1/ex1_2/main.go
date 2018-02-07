// Modify the echo3 program to print the index and value
// One per line
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " \n"))
}
