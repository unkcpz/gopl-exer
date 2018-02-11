// tempconv convert different temprature representing
package tempconv

import "fmt"

const (
	BoilingC   Celsius = 100.0
	FreezingC  Celsius = 0.0
	AbsolutedC Celsius = -273.15
)

type Celsius float64
type Fehrenheit float64

func (c Celsius) String() string    { return fmt.Sprintf("%.2fC", c) }
func (f Fehrenheit) String() string { return fmt.Sprintf("%.2fF", f) }
