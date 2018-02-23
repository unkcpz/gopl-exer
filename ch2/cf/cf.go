// cf converts arguments to Celsius and Ferenheit
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/unkcpz/gopl-exer/ch2/tempconv/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Println(err)
			continue
		}
		ct := tempconv.Celsius(t)
		ft := tempconv.Fehrenheit(t)
		fmt.Printf("%v = %v, %v = %v\n", ct, tempconv.CToF(ct), ft, tempconv.FToC(ft))
	}
}
