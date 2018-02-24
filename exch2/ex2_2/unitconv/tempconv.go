package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64

// String function for print
func (c Celsius) String() string {
	return fmt.Sprintf("%fC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%fF", f)
}

// CToF converts Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit((c + 32) * 9 / 5)
}

// KToC converts Kelvin to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius(f*5/9 - 32)
}
