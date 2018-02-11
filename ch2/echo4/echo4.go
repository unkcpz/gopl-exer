// Echo prints arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var n = flag.Bool("n", false, "new line at end")
	var sep = flag.String("s", " ", "separater betwenn words")
	flag.Parse()
	fmt.Printf("%s", strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Print("\n")
	}
}
