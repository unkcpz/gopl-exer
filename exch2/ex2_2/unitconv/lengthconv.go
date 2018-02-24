package unitconv

import "fmt"

type Meters float64
type Feet float64

func (m Meters) String() string {
	return fmt.Sprintf("%fM", m)
}

func (f Feet) String() string {
	return fmt.Sprintf("%fft in", f)
}

// MToF converts Meters to Feet
func MToF(m Meters) Feet {
	return Feet(m * 3.28084)
}

//FToM converts Feet to Meters
func FToM(f Feet) Meters {
	return Meters(f / 3.28084)
}
