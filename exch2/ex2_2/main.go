// General unit converts program
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/unkcpz/mygopl/exch2/ex2_2/unitconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 0 {
		for _, arg := range args {
			n, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				log.Println(err)
				continue
			}
			unitPrint(n)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			n, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				log.Println(err)
				continue
			}
			unitPrint(n)
		}
	}
}

func unitPrint(n float64) {
	ct := unitconv.Celsius(n)
	ft := unitconv.Fahrenheit(n)
	pw := unitconv.Pounds(n)
	kw := unitconv.Kilograms(n)
	ml := unitconv.Meters(n)
	fl := unitconv.Feet(n)
	fmt.Printf("%v = %v, %v = %v\n", ct, unitconv.CToF(ct), ft, unitconv.FToC(ft))
	fmt.Printf("%v = %v, %v = %v\n", pw, unitconv.PToK(pw), kw, unitconv.KToP(kw))
	fmt.Printf("%v = %v, %v = %v\n", ml, unitconv.MToF(ml), fl, unitconv.FToM(fl))
}
