// tempconv package converts different temprature repretension
package tempconv

import "fmt"

type Celsius float64
type Ferenheit float64
type Kelvin float64

const (
	AbsoluteZeroK = Kelvin(0.0)
	BoilingC      = Celsius(100.0)
	FreezingC     = Celsius(0.0)
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Ferenheit) String() string {
	return fmt.Sprintf("%gF", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}

// FToC converts Ferenheit to Celsius
func FToC(f Ferenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// KToC converts Kelvin to Celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k + 273.15)
}

// CToF converts Celsius to Ferenheit
func CToF(c Celsius) Ferenheit {
	return Ferenheit((c + 32) * 9 / 5)
}

// CToK converts Celsius to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c - 273.15)
}

// FToK converts Ferenheit to Kelvin
func FToK(f Ferenheit) Kelvin {
	return CToK(FToC(f))
}

// KToF converts Kelvin to Ferenheit
func KToF(k Kelvin) Ferenheit {
	return CToF(KToC(k))
}
